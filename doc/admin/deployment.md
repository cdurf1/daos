# DAOS System Deployment

## Preflight Checklist

This section covers the preliminary setup required on the compute and
storage nodes before deploying DAOS.

### Time Synchronization

The DAOS transaction model relies on timestamps and requires time to be
synchronized across all the storage and client nodes. This can be done
using NTP or any other equivalent protocol.

### Runtime Directory Setup

DAOS uses a series of Unix Domain Sockets to communicate between its
various components. On modern Linux systems, Unix Domain Sockets are
typically stored under /run or /var/run (usually a symlink to /run) and
are a mounted tmpfs file system. There are several methods for ensuring
the necessary directories are setup.

A sign that this step may have been missed is when starting daos_server
or daos_agent, you may see the message:
```
$ mkdir /var/run/daos_server: permission denied
Unable to create socket directory: /var/run/daos_server
```
#### Non-default Directory

By default, daos_server and daos_agent will use the directories
/var/run/daos_server and /var/run/daos_agent respectively. To change
the default location that daos_server uses for its runtime directory,
either uncomment and set the socket_dir configuration value in
install/etc/daos_server.yml, or pass the location to daos_server on
the command line using the -d flag. For the daos_agent, an alternate
location can be passed on the command line using the --runtime_dir flag.

#### Default Directory (non-persistent)

Files and directories created in /run and /var/run only survive until
the next reboot. However, if reboots are infrequent, an easy solution
while still utilizing the default locations is to create the
required directories manually. To do this execute the following commands.

daos_server:
```
$ mkdir /var/run/daos_server
$ chmod 0755 /var/run/daos_server
$ chown user:user /var/run/daos_server (where user is the user you
    will run daos_server as)
```
daos_agent:
```
$ mkdir /var/run/daos_agent
$ chmod 0755 /var/run/daos_agent
$ chown user:user /var/run/daos_agent (where user is the user you
    will run daos_agent as)
```

#### Default Directory (persistent)

If the server hosting daos_server or daos_agent will be rebooted often,
systemd provides a persistent mechanism for creating the required
directories called tmpfiles.d. This mechanism will be required every
time the system is provisioned and requires a reboot to take effect.

To tell systemd to create the necessary directories for DAOS:

-   Copy the file utils/systemd/daosfiles.conf to /etc/tmpfiles.d\
    cp utils/systemd/daosfiles.conf /etc/tmpfiles.d

-   Modify the copied file to change the user and group fields
    (currently daos) to the user daos will be run as

-   Reboot the system, and the directories will be created automatically
    on all subsequent reboots.

### Elevated Privileges

DAOS employs a privileged helper binary (`daos_admin`) to perform tasks
that require elevated privileges on behalf of `daos_server`.

Due to limitations introduced by recent security fixes in the kernel,
DAOS I/O processes are also required to run with elevated privileges
in order to access NVMe SSDs through the SPDK framework.

This is a temporary requirement that will be mitigated by moving to
use VFIO driver with SPDK (requires IOMMU enabled) rather than UIO.

## Typical Workflow

DAOS Control Server ([`daos_server`](/src/control/server)) instances will
listen for requests from the management tool ([`dmg`](/src/control/cmd/dmg)),
enabling users to perform remote operations in parallel across
multiple storage nodes.
The first type of commands run after installation include network and
storage hardware provisioning and would typically be run from a login
node.

After `daos_server` instances have been started on each storage node
for the first time (using a parallel launcher), `dmg storage prepare`
will set DCPM storage into the necessary state for use with DAOS.
Then `dmg storage format` formats persistent storage devices
(specified in the server configuration file) on the storage nodes
and writes necessary metadata before starting DAOS I/O processes that
will operate across the fabric.

Typically an administrator will perform the following steps:

1. Create or copy server configuration file
2. Start DAOS Control Servers (one per host)
3. Prepare DCPM modules for use with DAOS (one-time only, requires reboot)
4. Scan storage to check device details and state
5. Update device identifiers in server configuration file
6. Format storage (DAOS I/O processes will be auto-started after format)
7. (Optional) Reset DAOS storage

Note that starting the DAOS Control Server instances can be performed
automatically on boot if start-up scripts are registered with systemd.

The following subsections will cover each step in more detail.

### Step 1. Create or copy server configuration file

Example configuration files can be found in the
[examples directory](https://github.com/daos-stack/daos/blob/master/utils/config/examples)
and description of the parameters in
[`daos_server.yml`](https://github.com/daos-stack/daos/blob/master/utils/config/daos_server.yml).

### Step 2. Start DAOS Control Servers

One instance of the `daos_server` process is to be started per
storage node (please complete the
[pre-flight checklist](#preflight-checklist) first).

This is typically performed using a parallel launcher, for the
examples in this document, `clush` will be used.

`clush -w wolf-[118-121,130-133] daos_server start -o utils/config/examples/daos_server_sockets.yml`
will launch `daos_server` on the specified hosts connecting to
the `port` parameter value specified in the server config file.

Note that the DAOS Control Servers can be started automatically on
boot with systemd. Practically any parallel launcher can be used
e.g. pdsh, orterun.

### Step 3. Prepare DCPM modules for use with DAOS

DCPM preparation is required once per DAOS installation and
requires the DAOS Control Servers to be running as root.

This step requires a reboot to enable DCPM resource allocation
changes to be read by BIOS.

DCPM preparation can be performed from the management tool
`dmg storage prepare --scm-only` or using the Control Server directly
`sudo daos_server storage prepare --scm-only`.

After running the command a reboot will be required, then the Control
Servers will then need to be started again and the command run for a
second time to expose the namespace device to be used by DAOS.
([more details](#scm-preparation))

Example usage:
- `dmg -l wolf-[118-121,130-133] -i storage prepare --scm-only`
after running, the user should be prompted for a reboot.
- `clush -w wolf-[118-121,130-133] reboot`
- `clush -w wolf-[118-121,130-133] daos_server start -o utils/config/examples/daos_server_sockets.yml`
- `dmg -l wolf-[118-121,130-133] -i storage prepare --scm-only`
after running, `/dev/pmemX` devices should be available on each of the hosts.

### Step 4. Scan Storage to check device details and state

- `dmg -l wolf-[118-121,130-133] -i storage scan [--verbose]`
will display details of NVMe and DCPM devices on the specified hosts that can be used with DAOS.
[more details](#storage-detection--selection)

Note that other storage query commands are also available,
`dmg storage --help` for listings.

### Step 5. Add device identifiers to server config file

Edit `$DAOS_SRC/utils/config/examples/daos_server_sockets.yml`
to update the server configuration file with the information
received from scan to instruct which devices to use.
The `servers` section is a list specifying details for each DAOS I/O
instance to be started on the host (currently a maximum of 2 per host
is imposed).
Devices with the same NUMA rating/node/socket should be colocated on
a single DAOS I/O instance where possible.
[more details](#server-configuration)

- `bdev_list` should be populated with NVMe PCI addresses
- `scm_list` should be populated with DCPM interleaved set namespaces
(e.g. `/dev/pmem1`)
- DAOS Control Servers will need to be restarted on all hosts after
updates to the server configuration file.
- Pick one host in the system and set `access_points` to host and
listening port (probably going to be the same as server config `port`
parameter).
This will be the host which bootstraps the DAOS management service
(MS).

To illustrate, assume a cluster with homogenous hardware
configurations that returns the following from scan for each host:

```
[daos@wolf-72 daos_m]$ dmg -l wolf-7[1-2] -i storage scan --verbose
wolf-7[1-2]:10001: connected
-------
wolf-7[1-2]
-------
SCM Namespace Socket ID Capacity
------------- --------- --------
pmem0         0         2.90TB
pmem1         1         2.90TB

NVMe PCI     Model                FW Revision Socket ID Capacity
--------     -----                ----------- --------- --------
0000:81:00.0 INTEL SSDPED1K750GA  E2010325    0         750.00GB
0000:87:00.0 INTEL SSDPEDMD016T4  8DV10171    0         1.56TB
0000:da:00.0 INTEL SSDPED1K750GA  E2010325    1         750.00GB
```

In this situation, the configuration file `servers` section could be
populated as follows:

```
<snip>
port: 10001
access_points: ["wolf-71:10001"] # <----- updated
<snip>
servers:
-
  targets: 8                # count of storage targets per each server
  first_core: 0             # offset of the first core for service xstreams
  nr_xs_helpers: 2          # count of offload/helper xstreams per target
  fabric_iface: eth0        # map to OFI_INTERFACE=eth0
  fabric_iface_port: 31416  # map to OFI_PORT=31416
  log_mask: ERR             # map to D_LOG_MASK=ERR
  log_file: /tmp/server.log # map to D_LOG_FILE=/tmp/server.log
  env_vars:                 # influence DAOS IO Server behaviour by setting env variables
  - DAOS_MD_CAP=1024
  - CRT_CTX_SHARE_ADDR=0
  - CRT_TIMEOUT=30
  - FI_SOCKETS_MAX_CONN_RETRY=1
  - FI_SOCKETS_CONN_TIMEOUT=2000
  scm_mount: /mnt/daos  # map to -s /mnt/daos
  scm_class: dcpm
  scm_list: [/dev/pmem0] # <----- updated
  bdev_class: nvme
  bdev_list: ["0000:87:00.0", "0000:81:00.0"]  # <----- updated
-
  targets: 8                # count of storage targets per each server
  first_core: 0             # offset of the first core for service xstreams
  nr_xs_helpers: 2          # count of offload/helper xstreams per target
  fabric_iface: eth0        # map to OFI_INTERFACE=eth0
  fabric_iface_port: 31416  # map to OFI_PORT=31416
  log_mask: ERR             # map to D_LOG_MASK=ERR
  log_file: /tmp/server.log # map to D_LOG_FILE=/tmp/server.log
  env_vars:                 # influence DAOS IO Server behaviour by setting env variables
  - DAOS_MD_CAP=1024
  - CRT_CTX_SHARE_ADDR=0
  - CRT_TIMEOUT=30
  - FI_SOCKETS_MAX_CONN_RETRY=1
  - FI_SOCKETS_CONN_TIMEOUT=2000
  scm_mount: /mnt/daos  # map to -s /mnt/daos
  scm_class: dcpm
  scm_list: [/dev/pmem1] # <----- updated
  bdev_class: nvme
  bdev_list: ["0000:da:00.0"]  # <----- updated
<end>
```

### Step 6. Format Storage (from any node)

- Assuming `daos_server` is running on all desired hosts (see step
#2) and is in `format` mode (`SCM format required` displayed on DAOS
Control Server stdout).
- `dmg -l wolf-[118-121,130-133] -i storage format` will create
necessary metadata on the SCM mount at the path specified in the
config file. The device specified in the config file (`scm_list`)
will be mounted at the given path.
The NVMe devices specified in the config file (`bdev_list`) will be
formatted using SPDK (Optane drives may take a while) and configuration
written to each SCM mount to enable DAOS I/O instances to use designated
SSDs through SPDK.
- [management tool details](/src/control/cmd/dmg/README.md#storage-format)
- [SCM specific details](/src/control/server/README.md#scm-format)
- [NVMe specific details](/src/control/server/README.md#nvme-format)

After successful format, DAOS Control Servers will start DAOS IO
instances that have been specified in the server config file.

Successful start-up is indicated by the following on stdout:
`DAOS I/O server (v0.8.0) process 433456 started on rank 1 with 8 target, 2 helper XS per target, firstcore 0, host wolf-72.wolf.hpdd.intel.com.`

### Step 7. (optional) Reset DAOS storage

To reset the DAOS metadata across all hosts the system must be reformatted.
First ensure all `daos_server` processes on all hosts have been
stopped, then for each SCM mount specified in the config file
(`scm_mount` in the `servers` section) umount and wipe FS signatures.

Example illustration with two IO instances specified in the config file:
- `clush -w wolf-[118-121,130-133] umount /mnt/daos1`
- `clush -w wolf-[118-121,130-133] umount /mnt/daos0`
- `clush -w wolf-[118-121,130-133] wipefs -a /dev/pmem1`
- `clush -w wolf-[118-121,130-133] wipefs -a /dev/pmem0`
- Then restart DAOS Control Servers (step #2) and format (step #6).

## Hardware Provisioning

### Storage Preparation

#### SCM Preparation

This section addresses how to verify that Optane DC Persistent Memory
(DCPM) is correctly installed on the storage nodes, and how to configure
it in interleaved mode to be used by DAOS in AppDirect mode.
Instructions for other types of SCM may be covered in the future.

Provisioning the SCM occurs by configuring DCPM modules in AppDirect memory regions
(interleaved mode) in groups of modules local to a specific socket (NUMA), and
resultant nvdimm namespaces are defined by a device identifier (e.g., /dev/pmem0).

DCPM can be configured and managed through the
[ipmctl](https://github.com/intel/ipmctl) library and associated tool. The
ipmctl command can be run as root and has detailed man pages and
help output (use "ipmctl help" to display it).

The list of NVDIMMs can be displayed as follows:

```
ipmctl show -dimm
```

| DimmID | Capacity  | HealthState | ActionRequired | LockState | FWVersion     |
| ------ | --------- | ----------- | -------------- | --------- | ------------- |
| 0x0001 | 502.5 GiB | Healthy     | 0              | Disabled  | 01.00.00.5127 |
| 0x0101 | 502.5 GiB | Healthy     | 0              | Disabled  | 01.00.00.5127 |
| 0x1001 | 502.5 GiB | Healthy     | 0              | Disabled  | 01.00.00.5127 |
| 0x1101 | 502.5 GiB | Healthy     | 0              | Disabled  | 01.00.00.5127 |

Moreover, DAOS requires DCPM to be configured in interleaved mode. A
storage subcommand (prepare --scm-only) can be used as a "command mode"
invocation of *daos_server* and must be run as root. SCM modules will
be configured into interleaved regions with memory mode set to
"AppDirect" mode with one set per socket (each module is assigned to a socket,
and reports this via its NUMA rating).

`sudo daos_server [<app_opts>] storage prepare [--scm-only|-s] [<cmd_opts>]`
The first time the command is run, the SCM AppDirect regions will be created as
resource allocations on any available DCPM modules (one region per NUMA
node/socket). The regions are activated after BIOS reads the new resource
allocations, and after initial completion the command prints a
message to ask for a reboot (the command will not initiate reboot itself).

'sudo daos_server storage prepare --scm-only' should be run for a second time after
system reboot to create the pmem kernel devices (/dev/pmemX
namespaces created on the new SCM regions).

One namespace per region is created, and each namespace may take up to a few
minutes to create. Details of the pmem devices will be displayed in JSON format
on command completion.

Example output from the initial call (with the SCM modules set to default MemoryMode):

```bash
Memory allocation goals for SCM will be changed and namespaces modified, this
will be a destructive operation.  ensure namespaces are unmounted and SCM is
otherwise unused.
A reboot is required to process new memory allocation goals.
```

Example output from the subsequent call (SCM modules configured to AppDirect
mode, and host rebooted):

```bash
Memory allocation goals for SCM will be changed and namespaces modified. This
will be a destructive operation. Ensure namespaces are unmounted and the SCM
is otherwise unused.
creating SCM namespace, may take a few minutes...
creating SCM namespace, may take a few minutes...
Persistent memory kernel devices:
[{UUID:5d2f2517-9217-4d7d-9c32-70731c9ac11e Blockdev:pmem1 Dev:namespace1.0 NumaNode:1} {UUID:2bfe6c40-f79a-4b8e-bddf-ba81d4427b9b Blockdev:pmem0 Dev:namespace0.0 NumaNode:0}]
```

`sudo daos_server [<app_opts>] storage prepare [--scm-only|-s] --reset [<cmd_opts>]`

All namespaces are disabled and destroyed. The SCM regions are removed by
resetting modules into "MemoryMode" through resource allocations.

Note that undefined behavior may result if the namespaces/pmem kernel
devices are mounted before running reset (as per the printed warning).

A subsequent reboot is required for BIOS to read the new resource
allocations.

Example output when resetting the SCM modules:

```bash
Memory allocation goals for SCM will be changed and namespaces modified, this
will be a destructive operation.  ensure namespaces are unmounted and SCM is
otherwise unused.
removing SCM namespace, may take a few minutes...
removing SCM namespace, may take a few minutes...
resetting SCM memory allocations
A reboot is required to process new memory allocation goals.
```

#### NVMe Preparation

DAOS supports only NVMe-capable SSDs that are accessed directly from
userspace through the SPDK library.

NVMe access through SPDK as an unprivileged user can be enabled by
running the example command
`sudo daos_server storage prepare --nvme-only -p 4096 -u bob`.

This will perform the required setup in order for `daos_server` to be run
by user "bob" who will own the hugepage mountpoint directory and vfio
groups as needed in SPDK operations.

If the target-user is unspecified (`-u` short option), the target user
will be the issuer of the sudo command (or root if not using sudo).

The specification of hugepages (`-p` short option) defines the number
of huge pages to allocate for use by SPDK.

A list of PCI addresses can also be supplied to avoid unbinding all
PCI devices from the kernel, using the `-w` / `--pci-whitelist` option.

The `sudo daos_server [<app_opts>] storage prepare [--nvme-only|-n] [<cmd_opts>]`
command wraps the SPDK setup script to unbind the devices from
original kernel drivers and then binds the devices to a UIO driver
through which SPDK can communicate.

When a PCI address whitelist is not specified, SPDK access to all SSDs
will be enabled for the user (either the user executing sudo, the user
specified as --target-user, or effective user - in that order of precedence)
involving changing the ownership of relevant files in addition to SPDK setup.

The devices can then be bound back to the original drivers with the command
`sudo daos_server [<app_opts>] storage prepare [--nvme-only|-n] --reset [<cmd_opts>]`.

### Storage Detection & Selection

While the DAOS server auto-detects all the usable
storage, the administrator will still be provided with the ability through
the configuration file (see next section) to whitelist or blacklist the
storage devices to be (or not) used. This section covers how to manually
detect the storage devices potentially usable by DAOS to
populate the configuration file when the administrator wants to have
finer control over the storage selection.

`dmg storage scan` can be run to query remote running `daos_server`
processes over the management network.

`sudo daos_server storage scan` can be used to query `daos_server`
directly (scans locally-attached SSDs and Intel Persistent Memory
Modules usable by DAOS).

```bash
[daos@wolf-72 daos_m]$ dmg -l wolf-7[1-2] -i storage scan --verbose
wolf-[71-72]:10001: connected
------------
wolf-[71-72]
------------
SCM Namespace Socket ID Capacity
------------- --------- --------
pmem0         0         2.90TB
pmem1         1         2.90TB

NVMe PCI     Model                FW Revision Socket ID Capacity
--------     -----                ----------- --------- --------
0000:81:00.0 INTEL SSDPED1K750GA  E2010325    1         750.00GB
0000:87:00.0 INTEL SSDPEDMD016T4  8DV10171    1         1.56TB
0000:da:00.0 INTEL SSDPED1K750GA  E2010325    1         750.00GB
```

The NVMe PCI field above is what should be used in the server
configuration file to identified NVMe SSDs.

Devices with the same NUMA node/socket should be used in the same per-server
section of the server configuration file for best performance.

### Network Interface Detection and Selection

To display the fabric interface, OFI provider and NUMA node
combinations detected on the DAOS server, use the following command:
```
$ daos_server network scan --all

        fabric_iface: ib0
        provider: ofi+psm2
        pinned_numa_node: 0


        fabric_iface: ib1
        provider: ofi+psm2
        pinned_numa_node: 1


        fabric_iface: ib0
        provider: ofi+verbs;ofi_rxm
        pinned_numa_node: 0


        fabric_iface: ib1
        provider: ofi+verbs;ofi_rxm
        pinned_numa_node: 1


        fabric_iface: ib0
        provider: ofi+verbs
        pinned_numa_node: 0


        fabric_iface: ib1
        provider: ofi+verbs
        pinned_numa_node: 1


        fabric_iface: ib0
        provider: ofi+sockets
        pinned_numa_node: 0


        fabric_iface: ib1
        provider: ofi+sockets
        pinned_numa_node: 1


        fabric_iface: eth0
        provider: ofi+sockets
        pinned_numa_node: 0


        fabric_iface: lo
        provider: ofi+sockets
        pinned_numa_node: 0
```
The network scan leverages data from libfabric.  Results are ordered from
highest performance at the top to lowest performance at the bottom of the list.
Once the fabric_iface and provider pair has been chosen, those items and the
pinned_numa_node may be inserted directly into the corresponding sections within
daos_server.yml.  Note that the provider is currently the same for all DAOS
IO server instances and is configured once in the server configuration.  The fabric_iface and pinned_numa_node are configured for each IO server instance.

A list of providers that may be querried is found with the command:
```
$ daos_server network list

Supported providers:
        ofi+gni, ofi+psm2, ofi+tcp, ofi+sockets, ofi+verbs, ofi_rxm
```

Performing a network scan that filters on a specific provider is accomplished
by issuing the following command:
```
$ daos_server network scan --provider 'ofi+verbs;ofi_rxm'

Scanning fabric for cmdline specified provider: ofi+verbs;ofi_rxm
Fabric scan found 2 devices matching the provider spec: ofi+verbs;ofi_rxm

        fabric_iface: ib0
        provider: ofi+verbs;ofi_rxm
        pinned_numa_node: 0


        fabric_iface: ib1
        provider: ofi+verbs;ofi_rxm
        pinned_numa_node: 1
```
To aid in provider configuration and debug, it may be helpful to run the
fi_pingpong test (delivered as part of OFI/libfabric).  To run that test,
determine the name of the provider to test usually by removing the "ofi+" prefix from the network scan provider data.  Do use the "ofi+" prefix in the
daos_server.yml.  Do not use the "ofi+" prefix with fi_pingpong.

Then, the fi_pingpong test can be used to verify that the targeted OFI provider works fine:
```
node1$ fi_pingpong -p psm2

node2$ fi_pingpong -p psm2 ${IP_ADDRESS_NODE1}

bytes #sent #ack total time  MB/sec  usec/xfer Mxfers/sec
64    10    =10  1.2k  0.00s 21.69   2.95      0.34
256   10    =10  5k    0.00s 116.36  2.20      0.45
1k    10    =10  20k   0.00s 379.26  2.70      0.37
4k    10    =10  80k   0.00s 1077.89 3.80      0.26
64k   10    =10  1.2m  0.00s 2145.20 30.55     0.03
1m    10    =10  20m   0.00s 8867.45 118.25    0.01
```

## Server Configuration

This section addresses how to configure the DAOS servers on the storage
nodes before starting it.

### Certificate Generation

The DAOS security framework relies on certificates to authenticate
administrators. The security infrastructure is currently under
development and will be delivered in DAOS v1.0. Initial support for certificates
has been added to DAOS and can be disabled either via the command line or in the
DAOS server configuration file. Currently, the easiest way to disable certificate
support is to pass the -i flag to daos_server.

### Server Configuration File

The `daos_server` configuration file is parsed when starting the
`daos_server` process. The configuration file location can be specified
on the command line (`daos_server -h` for usage) or default location
(`install/etc/daos_server.yml`).

Parameter descriptions are specified in [`daos_server.yml`](https://github.com/daos-stack/daos/blob/master/utils/config/daos_server.yml)
and example configuration files in the [examples](https://github.com/daos-stack/daos/tree/master/utils/config/examples)
directory.

Any option supplied to `daos_server` as a command line option or flag will
take precedence over equivalent configuration file parameter.

For convenience, active parsed configuration values are written to a temporary
file for reference, and the location will be written to the log.

#### Configuration File Options

The example configuration file lists the default empty configuration, listing all the
options (living documentation of the config file). Live examples are
available at
<https://github.com/daos-stack/daos/tree/master/utils/config>

The location of this configuration file is determined by first checking
for the path specified through the -o option of the daos_server command
line. Otherwise, /etc/daos_server.conf is used.

Refer to the example configuration file ([daos_server.yml](https://github.com/daos-stack/daos/blob/master/utils/config/daos_server.yml))
for latest information and examples.

## Server Startup

DAOS is currently switching from the PMIx-based server wire-up to a
self-contained bootstrap procedure. The new bootstrap procedure will be
available for DAOS v1.0 and will allow the DAOS servers to be started
individually (e.g. independently on each storage node via systemd) or
collectively (e.g. pdsh, mpirun or as a Kubernetes Pod). Meanwhile, servers
no longer have to be started by orterun, with a temporary limitation that if
one of them is restarted, the others must also be restarted.

### Parallel Launcher

As stated above, orterun(1) is no longer required, provided the temporary
limitation is accommodated. The section still uses orterun as an example.

The list of storage nodes can be specified on the command line via the -H
option. The DAOS server and the application can be started separately.
Also, the DAOS server must be started with the --enable-recovery option
to support server failure. See the orterun(1) man page for additional options.

To start the DAOS server, run:
```
orterun --map-by node --mca btl tcp,self --mca oob tcp -np <num_servers>
-H <server_list> --enable-recovery daos_server start -o <config_file>
```
The --enable-recovery is required for fault tolerance to guarantee that
the fault of one server does not cause the others to be stopped.

The --allow-run-as-root option can be added to the command line to
allow the daos_server to run with root privileges on each storage
nodes (for example when needing to perform privileged tasks relating
to storage format).

The content of the configuration file is documented in the next section
and a few examples are [available](/src/utils/config/examples).

### Systemd Integration

Systemd support for daos_server is still experimental as it will start the
daos_server and daos_io_server components in PMIXless mode, which is still in
development.

DAOS Server can be started as a systemd service. The DAOS Server
unit file is installed in the correct location when installing from RPMs.
If you wish to use systemd with a development build, you must copy the service
file from utils/systemd to /usr/lib/systemd/system. Once the file is copied
modify the ExecStart line to point to your in tree daos_server binary.

Once the service file is installed you can start daos_server
with the following commands:

```bash
$ systemctl enable daos-server
$ systemctl start daos-server
```

To check the component status use:

```bash
$ systemctl status daos-server
```

If DAOS Server failed to start, check the logs with:

```bash
$ journalctl --unit daos-server
```

### Kubernetes Pod

DAOS service integration with Kubernetes is planned and will be
supported in a future DAOS version.

### Internal DAOS Control Server Components

On start-up, the `daos_server` will create and initialize the following
components:

- gRPC server to handle requests over client API

- dRPC server to handle requests from IO servers over the UNIX domain
socket

- communication primitives for request/response to/from privileged binary
helper `daos_admin`

## Storage Formatting

When `daos_server` is started for the first time (and no SCM directory exists),
it enters "maintenance mode" and waits for a `dmg storage format` call to
be issued from the management tool.
This remote call will trigger the formatting of the locally attached storage on
the host for use with DAOS using the parameters defined in the server config file.

`dmg -i -l <host:port>[,...] storage format` will normally be run on a login
node specifying a hostlist (`-l <host:port>[,...]`) of storage nodes with SCM/DCPM
modules and NVMe SSDs installed and prepared.

### SCM Format

When the command is run, the pmem kernel devices created on SCM/DCPM regions are
formatted and mounted based on the parameters provided in the server config file.

- `scm_mount` specifies the location of the mountpoint to create.
- `scm_class` can be set to `ram` to use a tmpfs in the situation that no SCM/DCPM
is available (scm_size dictates the size of tmpfs in GB), when set to `dcpm` the device
specified under `scm_list` will be mounted at `scm_mount` path.

### NVMe Format

When the command is run, NVMe SSDs are formatted and set up to be used by DAOS
based on the parameters provided in the server config file.

`bdev_class` can be set to `nvme` to use actual NVMe devices with SPDK for DAOS
storage.
Other `bdev_class` values can be used for emulation of NVMe storage as specified
in the server config file.
`bdev_list` identifies devices to use with a list of PCI addresses (this can be
populated after viewing results from `storage scan` command).

After the format command is run, the path specified by the server configuration
file `scm_mount` parameter should be mounted and should contain a file named
`daos_nvme.conf`.
The file should describe the devices with PCI addresses as listed in the
`bdev_list` parameter of the server config file.
The presence and contents of the file indicate that the specified NVMe SSDs have
been configured correctly for use with DAOS.

The contents of the NVMe SSDs listed in the server configuration file `bdev_list`
parameter will be reset on format.

### Server Format

Before the format command is run, no DAOS metadata should exist under the
path specified by `scm_mount` parameter in the server configuration file.

After the `storage format` command is run, the path specified by the server
configuration file `scm_mount` parameter should be mounted and should contain
the necessary DAOS metadata indicating that the server has been formatted.

When starting, `daos_server` will skip `maintenance mode` and attempt to start
IO services if valid DAOS metadata is found in `scm_mount`.

## Agent Configuration

This section addresses how to configure the DAOS agents on the storage
nodes before starting it.

### Agent Certificate Generation

The DAOS security framework relies on certificates to authenticate
administrators. The security infrastructure is currently under
development and will be delivered in DAOS v1.0. Initial support for certificates
has been added to DAOS and can be disabled either via the command line or in the
DAOS Agent configuration file. Currently, the easiest way to disable certificate
support is to pass the -i flag to daos_agent.

### Agent Configuration File

The `daos_agent` configuration file is parsed when starting the
`daos_agent` process. The configuration file location can be specified
on the command line (`daos_agent -h` for usage) or default location
(`install/etc/daos_agent.yml`).

Parameter descriptions are specified in [daos_agent.yml](https://github.com/daos-stack/daos/blob/master/utils/config/daos_agent.yml).

Any option supplied to `daos_agent` as a command line option or flag will
take precedence over equivalent configuration file parameter.

For convenience, active parsed config values are written to a temporary
file for reference, and the location will be written to the log.

The following section lists the format, options, defaults, and
descriptions available in the configuration file.

#### Configuration File Options

The example configuration file lists the default empty configuration listing all the
options (living documentation of the config file). Live examples are
available at
<https://github.com/daos-stack/daos/tree/master/utils/config>

The location of this configuration file is determined by first checking
for the path specified through the -o option of the daos_agent command
line. Otherwise, /etc/daos_agent.conf is used.

Refer to the example configuration file ([daos_server.yml](https://github.com/daos-stack/daos/blob/master/utils/config/daos_server.yml)) for latest information and examples.

## Agent Startup

DAOS Agent is a standalone application to be run on each compute node.
It can be configured to use secure communications (default) or can be allowed
to communicate with the control plane over unencrypted channels. The following
example shows daos_agent being configured to operate in insecure mode due to
incomplete integration of certificate support as of the 0.6 release.

To start the DAOS Agent from the command line, run:

```bash
$ daos_agent -i
```

Alternatively, the DAOS Agent can be started as a systemd service. The DAOS Agent
unit file is installed in the correct location when installing from RPMs.
If you wish to use systemd with a development build, you must copy the service
file from utils/systemd to /usr/lib/systemd/system. Once the file is copied
modify the ExecStart line to point to your in tree daos_agent binary.

Once the service file is installed, you can start daos_agent
with the following commands:

```bash
$ systemctl enable daos-agent
$ systemctl start daos-agent
```

To check the component status use:

```bash
$ systemctl status daos-agent
```

If DAOS Agent failed to start check the logs with:

```bash
$ journalctl --unit daos-agent
```

## System Validation

To validate that the DAOS system is properly installed, the daos_test
suite can be executed. Ensure the DAOS Agent is configured and running before
running daos_test and that the DAOS_SINGLETON_CLI and CRT_ATTACH_INFO_PATH
environment variables are properly set as described [here](#server-startup).

```
orterun -np <num_clients> --hostfile <hostfile> ./daos_test
```

daos_test requires at least 8GB of SCM (or DRAM with tmpfs) storage on
each storage node.

[^1]: https://github.com/intel/ipmctl

[^2]: https://github.com/daos-stack/daos/tree/master/utils/config

[^3]: [*https://www.open-mpi.org/faq/?category=running\#mpirun-hostfile*](https://www.open-mpi.org/faq/?category=running#mpirun-hostfile)

[^4]: https://github.com/daos-stack/daos/tree/master/src/control/README.md
