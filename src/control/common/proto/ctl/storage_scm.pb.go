// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_scm.proto

package ctl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ScmModule represent Storage Class Memory modules installed.
type ScmModule struct {
	//string uid = 1; // The uid of the module.
	Physicalid uint32 `protobuf:"varint,1,opt,name=physicalid,proto3" json:"physicalid,omitempty"`
	//string handle = 3; // The device handle of the module.
	//string serial = 8; // The serial number of the module.
	Capacity uint64 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	//string fwrev = 10; // The firmware revision of the module.
	Loc                  *ScmModule_Location `protobuf:"bytes,3,opt,name=loc,proto3" json:"loc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ScmModule) Reset()         { *m = ScmModule{} }
func (m *ScmModule) String() string { return proto.CompactTextString(m) }
func (*ScmModule) ProtoMessage()    {}
func (*ScmModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{0}
}

func (m *ScmModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModule.Unmarshal(m, b)
}
func (m *ScmModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModule.Marshal(b, m, deterministic)
}
func (m *ScmModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModule.Merge(m, src)
}
func (m *ScmModule) XXX_Size() int {
	return xxx_messageInfo_ScmModule.Size(m)
}
func (m *ScmModule) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModule.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModule proto.InternalMessageInfo

func (m *ScmModule) GetPhysicalid() uint32 {
	if m != nil {
		return m.Physicalid
	}
	return 0
}

func (m *ScmModule) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *ScmModule) GetLoc() *ScmModule_Location {
	if m != nil {
		return m.Loc
	}
	return nil
}

type ScmModule_Location struct {
	Channel              uint32   `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Channelpos           uint32   `protobuf:"varint,2,opt,name=channelpos,proto3" json:"channelpos,omitempty"`
	Memctrlr             uint32   `protobuf:"varint,3,opt,name=memctrlr,proto3" json:"memctrlr,omitempty"`
	Socket               uint32   `protobuf:"varint,4,opt,name=socket,proto3" json:"socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScmModule_Location) Reset()         { *m = ScmModule_Location{} }
func (m *ScmModule_Location) String() string { return proto.CompactTextString(m) }
func (*ScmModule_Location) ProtoMessage()    {}
func (*ScmModule_Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{0, 0}
}

func (m *ScmModule_Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModule_Location.Unmarshal(m, b)
}
func (m *ScmModule_Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModule_Location.Marshal(b, m, deterministic)
}
func (m *ScmModule_Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModule_Location.Merge(m, src)
}
func (m *ScmModule_Location) XXX_Size() int {
	return xxx_messageInfo_ScmModule_Location.Size(m)
}
func (m *ScmModule_Location) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModule_Location.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModule_Location proto.InternalMessageInfo

func (m *ScmModule_Location) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *ScmModule_Location) GetChannelpos() uint32 {
	if m != nil {
		return m.Channelpos
	}
	return 0
}

func (m *ScmModule_Location) GetMemctrlr() uint32 {
	if m != nil {
		return m.Memctrlr
	}
	return 0
}

func (m *ScmModule_Location) GetSocket() uint32 {
	if m != nil {
		return m.Socket
	}
	return 0
}

// PmemDevice represents SCM namespace as pmem device files created on a ScmRegion.
type PmemDevice struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Blockdev             string   `protobuf:"bytes,2,opt,name=blockdev,proto3" json:"blockdev,omitempty"`
	Dev                  string   `protobuf:"bytes,3,opt,name=dev,proto3" json:"dev,omitempty"`
	Numanode             uint32   `protobuf:"varint,4,opt,name=numanode,proto3" json:"numanode,omitempty"`
	Size                 uint64   `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmemDevice) Reset()         { *m = PmemDevice{} }
func (m *PmemDevice) String() string { return proto.CompactTextString(m) }
func (*PmemDevice) ProtoMessage()    {}
func (*PmemDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{1}
}

func (m *PmemDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmemDevice.Unmarshal(m, b)
}
func (m *PmemDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmemDevice.Marshal(b, m, deterministic)
}
func (m *PmemDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmemDevice.Merge(m, src)
}
func (m *PmemDevice) XXX_Size() int {
	return xxx_messageInfo_PmemDevice.Size(m)
}
func (m *PmemDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_PmemDevice.DiscardUnknown(m)
}

var xxx_messageInfo_PmemDevice proto.InternalMessageInfo

func (m *PmemDevice) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PmemDevice) GetBlockdev() string {
	if m != nil {
		return m.Blockdev
	}
	return ""
}

func (m *PmemDevice) GetDev() string {
	if m != nil {
		return m.Dev
	}
	return ""
}

func (m *PmemDevice) GetNumanode() uint32 {
	if m != nil {
		return m.Numanode
	}
	return 0
}

func (m *PmemDevice) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// ScmMount represents mounted AppDirect region made up of SCM module set.
type ScmMount struct {
	Mntpoint             string       `protobuf:"bytes,1,opt,name=mntpoint,proto3" json:"mntpoint,omitempty"`
	Modules              []*ScmModule `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
	Pmem                 *PmemDevice  `protobuf:"bytes,3,opt,name=pmem,proto3" json:"pmem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScmMount) Reset()         { *m = ScmMount{} }
func (m *ScmMount) String() string { return proto.CompactTextString(m) }
func (*ScmMount) ProtoMessage()    {}
func (*ScmMount) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{2}
}

func (m *ScmMount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmMount.Unmarshal(m, b)
}
func (m *ScmMount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmMount.Marshal(b, m, deterministic)
}
func (m *ScmMount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmMount.Merge(m, src)
}
func (m *ScmMount) XXX_Size() int {
	return xxx_messageInfo_ScmMount.Size(m)
}
func (m *ScmMount) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmMount.DiscardUnknown(m)
}

var xxx_messageInfo_ScmMount proto.InternalMessageInfo

func (m *ScmMount) GetMntpoint() string {
	if m != nil {
		return m.Mntpoint
	}
	return ""
}

func (m *ScmMount) GetModules() []*ScmModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *ScmMount) GetPmem() *PmemDevice {
	if m != nil {
		return m.Pmem
	}
	return nil
}

// ScmModuleResult represents operation state for specific SCM/PM module.
//
// TODO: replace identifier with serial when returned in scan
type ScmModuleResult struct {
	Loc                  *ScmModule_Location `protobuf:"bytes,1,opt,name=loc,proto3" json:"loc,omitempty"`
	State                *ResponseState      `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ScmModuleResult) Reset()         { *m = ScmModuleResult{} }
func (m *ScmModuleResult) String() string { return proto.CompactTextString(m) }
func (*ScmModuleResult) ProtoMessage()    {}
func (*ScmModuleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{3}
}

func (m *ScmModuleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModuleResult.Unmarshal(m, b)
}
func (m *ScmModuleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModuleResult.Marshal(b, m, deterministic)
}
func (m *ScmModuleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModuleResult.Merge(m, src)
}
func (m *ScmModuleResult) XXX_Size() int {
	return xxx_messageInfo_ScmModuleResult.Size(m)
}
func (m *ScmModuleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModuleResult.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModuleResult proto.InternalMessageInfo

func (m *ScmModuleResult) GetLoc() *ScmModule_Location {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *ScmModuleResult) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

// ScmMountResult represents operation state for specific SCM mount point.
type ScmMountResult struct {
	Mntpoint             string         `protobuf:"bytes,1,opt,name=mntpoint,proto3" json:"mntpoint,omitempty"`
	State                *ResponseState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ScmMountResult) Reset()         { *m = ScmMountResult{} }
func (m *ScmMountResult) String() string { return proto.CompactTextString(m) }
func (*ScmMountResult) ProtoMessage()    {}
func (*ScmMountResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{4}
}

func (m *ScmMountResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmMountResult.Unmarshal(m, b)
}
func (m *ScmMountResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmMountResult.Marshal(b, m, deterministic)
}
func (m *ScmMountResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmMountResult.Merge(m, src)
}
func (m *ScmMountResult) XXX_Size() int {
	return xxx_messageInfo_ScmMountResult.Size(m)
}
func (m *ScmMountResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmMountResult.DiscardUnknown(m)
}

var xxx_messageInfo_ScmMountResult proto.InternalMessageInfo

func (m *ScmMountResult) GetMntpoint() string {
	if m != nil {
		return m.Mntpoint
	}
	return ""
}

func (m *ScmMountResult) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

type PrepareScmReq struct {
	Reset_               bool     `protobuf:"varint,1,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareScmReq) Reset()         { *m = PrepareScmReq{} }
func (m *PrepareScmReq) String() string { return proto.CompactTextString(m) }
func (*PrepareScmReq) ProtoMessage()    {}
func (*PrepareScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{5}
}

func (m *PrepareScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareScmReq.Unmarshal(m, b)
}
func (m *PrepareScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareScmReq.Marshal(b, m, deterministic)
}
func (m *PrepareScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareScmReq.Merge(m, src)
}
func (m *PrepareScmReq) XXX_Size() int {
	return xxx_messageInfo_PrepareScmReq.Size(m)
}
func (m *PrepareScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareScmReq proto.InternalMessageInfo

func (m *PrepareScmReq) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type PrepareScmResp struct {
	Pmems                []*PmemDevice  `protobuf:"bytes,1,rep,name=pmems,proto3" json:"pmems,omitempty"`
	State                *ResponseState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PrepareScmResp) Reset()         { *m = PrepareScmResp{} }
func (m *PrepareScmResp) String() string { return proto.CompactTextString(m) }
func (*PrepareScmResp) ProtoMessage()    {}
func (*PrepareScmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{6}
}

func (m *PrepareScmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareScmResp.Unmarshal(m, b)
}
func (m *PrepareScmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareScmResp.Marshal(b, m, deterministic)
}
func (m *PrepareScmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareScmResp.Merge(m, src)
}
func (m *PrepareScmResp) XXX_Size() int {
	return xxx_messageInfo_PrepareScmResp.Size(m)
}
func (m *PrepareScmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareScmResp.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareScmResp proto.InternalMessageInfo

func (m *PrepareScmResp) GetPmems() []*PmemDevice {
	if m != nil {
		return m.Pmems
	}
	return nil
}

func (m *PrepareScmResp) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

type ScanScmReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanScmReq) Reset()         { *m = ScanScmReq{} }
func (m *ScanScmReq) String() string { return proto.CompactTextString(m) }
func (*ScanScmReq) ProtoMessage()    {}
func (*ScanScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{7}
}

func (m *ScanScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanScmReq.Unmarshal(m, b)
}
func (m *ScanScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanScmReq.Marshal(b, m, deterministic)
}
func (m *ScanScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanScmReq.Merge(m, src)
}
func (m *ScanScmReq) XXX_Size() int {
	return xxx_messageInfo_ScanScmReq.Size(m)
}
func (m *ScanScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_ScanScmReq proto.InternalMessageInfo

type ScanScmResp struct {
	Modules              []*ScmModule   `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	Pmems                []*PmemDevice  `protobuf:"bytes,2,rep,name=pmems,proto3" json:"pmems,omitempty"`
	State                *ResponseState `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ScanScmResp) Reset()         { *m = ScanScmResp{} }
func (m *ScanScmResp) String() string { return proto.CompactTextString(m) }
func (*ScanScmResp) ProtoMessage()    {}
func (*ScanScmResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{8}
}

func (m *ScanScmResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanScmResp.Unmarshal(m, b)
}
func (m *ScanScmResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanScmResp.Marshal(b, m, deterministic)
}
func (m *ScanScmResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanScmResp.Merge(m, src)
}
func (m *ScanScmResp) XXX_Size() int {
	return xxx_messageInfo_ScanScmResp.Size(m)
}
func (m *ScanScmResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanScmResp.DiscardUnknown(m)
}

var xxx_messageInfo_ScanScmResp proto.InternalMessageInfo

func (m *ScanScmResp) GetModules() []*ScmModule {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *ScanScmResp) GetPmems() []*PmemDevice {
	if m != nil {
		return m.Pmems
	}
	return nil
}

func (m *ScanScmResp) GetState() *ResponseState {
	if m != nil {
		return m.State
	}
	return nil
}

type FormatScmReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormatScmReq) Reset()         { *m = FormatScmReq{} }
func (m *FormatScmReq) String() string { return proto.CompactTextString(m) }
func (*FormatScmReq) ProtoMessage()    {}
func (*FormatScmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa79a1cba4dc284c, []int{9}
}

func (m *FormatScmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormatScmReq.Unmarshal(m, b)
}
func (m *FormatScmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormatScmReq.Marshal(b, m, deterministic)
}
func (m *FormatScmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormatScmReq.Merge(m, src)
}
func (m *FormatScmReq) XXX_Size() int {
	return xxx_messageInfo_FormatScmReq.Size(m)
}
func (m *FormatScmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FormatScmReq.DiscardUnknown(m)
}

var xxx_messageInfo_FormatScmReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScmModule)(nil), "ctl.ScmModule")
	proto.RegisterType((*ScmModule_Location)(nil), "ctl.ScmModule.Location")
	proto.RegisterType((*PmemDevice)(nil), "ctl.PmemDevice")
	proto.RegisterType((*ScmMount)(nil), "ctl.ScmMount")
	proto.RegisterType((*ScmModuleResult)(nil), "ctl.ScmModuleResult")
	proto.RegisterType((*ScmMountResult)(nil), "ctl.ScmMountResult")
	proto.RegisterType((*PrepareScmReq)(nil), "ctl.PrepareScmReq")
	proto.RegisterType((*PrepareScmResp)(nil), "ctl.PrepareScmResp")
	proto.RegisterType((*ScanScmReq)(nil), "ctl.ScanScmReq")
	proto.RegisterType((*ScanScmResp)(nil), "ctl.ScanScmResp")
	proto.RegisterType((*FormatScmReq)(nil), "ctl.FormatScmReq")
}

func init() { proto.RegisterFile("storage_scm.proto", fileDescriptor_fa79a1cba4dc284c) }

var fileDescriptor_fa79a1cba4dc284c = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x6e, 0x13, 0x41,
	0x10, 0xd4, 0x64, 0xed, 0xc4, 0x6e, 0x3f, 0x02, 0x23, 0x04, 0x2b, 0x1f, 0x90, 0xb5, 0x28, 0xd2,
	0x72, 0xf1, 0xc1, 0xfc, 0x02, 0xe2, 0x04, 0x52, 0x34, 0x96, 0xb8, 0xa2, 0xc9, 0xb8, 0x21, 0xab,
	0xcc, 0x8b, 0x9d, 0xd9, 0x88, 0x70, 0xe0, 0xcc, 0xc7, 0xf2, 0x11, 0x68, 0x7a, 0x1f, 0x31, 0x42,
	0x4a, 0x9c, 0xdb, 0x54, 0x77, 0x6d, 0x57, 0x75, 0xb5, 0x16, 0x9e, 0x87, 0xe8, 0x6a, 0xf9, 0x0d,
	0xbf, 0x04, 0x65, 0x36, 0xbe, 0x76, 0xd1, 0xf1, 0x4c, 0x45, 0xbd, 0x9a, 0x2b, 0x67, 0x8c, 0xb3,
	0x6d, 0xa9, 0xf8, 0xc3, 0x60, 0xba, 0x53, 0xe6, 0x93, 0xdb, 0x37, 0x1a, 0xf9, 0x6b, 0x00, 0x7f,
	0x7d, 0x17, 0x2a, 0x25, 0x75, 0xb5, 0xcf, 0xd9, 0x9a, 0x95, 0x0b, 0x71, 0x50, 0xe1, 0x2b, 0x98,
	0x28, 0xe9, 0xa5, 0xaa, 0xe2, 0x5d, 0x7e, 0xb2, 0x66, 0xe5, 0x48, 0x0c, 0x98, 0xbf, 0x85, 0x4c,
	0x3b, 0x95, 0x67, 0x6b, 0x56, 0xce, 0xb6, 0xaf, 0x36, 0x2a, 0xea, 0xcd, 0x30, 0x78, 0xf3, 0xd1,
	0x29, 0x19, 0x2b, 0x67, 0x45, 0xe2, 0xac, 0x7e, 0xc0, 0xa4, 0x2f, 0xf0, 0x1c, 0xce, 0xd4, 0xb5,
	0xb4, 0x16, 0x75, 0xa7, 0xd7, 0xc3, 0x64, 0xa6, 0x7b, 0x7a, 0x17, 0x48, 0x6e, 0x21, 0x0e, 0x2a,
	0xc9, 0x8c, 0x41, 0xa3, 0x62, 0xad, 0x6b, 0x52, 0x5d, 0x88, 0x01, 0xf3, 0x97, 0x70, 0x1a, 0x9c,
	0xba, 0xc1, 0x98, 0x8f, 0xa8, 0xd3, 0xa1, 0xe2, 0x17, 0xc0, 0xa5, 0x41, 0xf3, 0x1e, 0x6f, 0x2b,
	0x85, 0x9c, 0xc3, 0xa8, 0x69, 0xba, 0x45, 0xa7, 0x82, 0xde, 0x69, 0xea, 0x95, 0x76, 0xea, 0x66,
	0x8f, 0xb7, 0xa4, 0x39, 0x15, 0x03, 0xe6, 0xcf, 0x20, 0x4b, 0xe5, 0x8c, 0xca, 0xe9, 0x99, 0xd8,
	0xb6, 0x31, 0xd2, 0xba, 0x3d, 0x76, 0x4a, 0x03, 0x4e, 0xd3, 0x43, 0xf5, 0x13, 0xf3, 0x31, 0x05,
	0x45, 0xef, 0xa2, 0x81, 0x09, 0x85, 0xd2, 0xd8, 0x48, 0xfe, 0x6d, 0xf4, 0xae, 0xb2, 0xb1, 0x73,
	0x30, 0x60, 0x5e, 0xc2, 0x99, 0xa1, 0xe4, 0xd2, 0xe2, 0x59, 0x39, 0xdb, 0x2e, 0xff, 0x0d, 0x54,
	0xf4, 0x6d, 0xfe, 0x06, 0x46, 0xde, 0xa0, 0xe9, 0x72, 0x3f, 0x27, 0xda, 0xfd, 0x8a, 0x82, 0x9a,
	0xc5, 0x57, 0x38, 0xbf, 0xff, 0x14, 0x43, 0xa3, 0x63, 0x7f, 0x2e, 0xf6, 0xf8, 0xb9, 0x78, 0x09,
	0xe3, 0x10, 0x65, 0x44, 0xca, 0x63, 0xb6, 0xe5, 0x44, 0x16, 0x18, 0xbc, 0xb3, 0x01, 0x77, 0xa9,
	0x23, 0x5a, 0x42, 0xf1, 0x19, 0x96, 0xfd, 0x7a, 0x9d, 0xcc, 0xc3, 0x4b, 0x1e, 0x3b, 0xf7, 0x02,
	0x16, 0x97, 0x35, 0x7a, 0x59, 0xe3, 0x4e, 0x19, 0x81, 0xdf, 0xf9, 0x0b, 0x18, 0xd7, 0x18, 0xb0,
	0x9d, 0x39, 0x11, 0x2d, 0x28, 0x24, 0x2c, 0x0f, 0x69, 0xc1, 0xf3, 0x0b, 0x18, 0xa7, 0x00, 0x42,
	0xce, 0x28, 0xc5, 0xff, 0xe2, 0x69, 0xbb, 0x4f, 0x70, 0x32, 0x07, 0xd8, 0x29, 0x69, 0x5b, 0x1b,
	0xc5, 0x6f, 0x06, 0xb3, 0x01, 0x06, 0x7f, 0x78, 0x36, 0xf6, 0xf0, 0xd9, 0x06, 0x63, 0x27, 0xc7,
	0x19, 0xcb, 0x1e, 0x33, 0xb6, 0x84, 0xf9, 0x07, 0x57, 0x1b, 0x19, 0x5b, 0x6b, 0x57, 0xa7, 0xf4,
	0x7f, 0xbf, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x19, 0x36, 0x0c, 0x96, 0x07, 0x04, 0x00, 0x00,
}
