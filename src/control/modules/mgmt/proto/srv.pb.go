// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListFeaturesParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFeaturesParams) Reset()         { *m = ListFeaturesParams{} }
func (m *ListFeaturesParams) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesParams) ProtoMessage()    {}
func (*ListFeaturesParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{0}
}
func (m *ListFeaturesParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeaturesParams.Unmarshal(m, b)
}
func (m *ListFeaturesParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeaturesParams.Marshal(b, m, deterministic)
}
func (dst *ListFeaturesParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeaturesParams.Merge(dst, src)
}
func (m *ListFeaturesParams) XXX_Size() int {
	return xxx_messageInfo_ListFeaturesParams.Size(m)
}
func (m *ListFeaturesParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeaturesParams.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeaturesParams proto.InternalMessageInfo

type FeatureName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureName) Reset()         { *m = FeatureName{} }
func (m *FeatureName) String() string { return proto.CompactTextString(m) }
func (*FeatureName) ProtoMessage()    {}
func (*FeatureName) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{1}
}
func (m *FeatureName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureName.Unmarshal(m, b)
}
func (m *FeatureName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureName.Marshal(b, m, deterministic)
}
func (dst *FeatureName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureName.Merge(dst, src)
}
func (m *FeatureName) XXX_Size() int {
	return xxx_messageInfo_FeatureName.Size(m)
}
func (m *FeatureName) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureName.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureName proto.InternalMessageInfo

func (m *FeatureName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Feature struct {
	// The name of the feature.
	Fname *FeatureName `protobuf:"bytes,1,opt,name=fname,proto3" json:"fname,omitempty"`
	// The description of the feature.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{2}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetFname() *FeatureName {
	if m != nil {
		return m.Fname
	}
	return nil
}

func (m *Feature) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ListNVMeParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNVMeParams) Reset()         { *m = ListNVMeParams{} }
func (m *ListNVMeParams) String() string { return proto.CompactTextString(m) }
func (*ListNVMeParams) ProtoMessage()    {}
func (*ListNVMeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{3}
}
func (m *ListNVMeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNVMeParams.Unmarshal(m, b)
}
func (m *ListNVMeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNVMeParams.Marshal(b, m, deterministic)
}
func (dst *ListNVMeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNVMeParams.Merge(dst, src)
}
func (m *ListNVMeParams) XXX_Size() int {
	return xxx_messageInfo_ListNVMeParams.Size(m)
}
func (m *ListNVMeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNVMeParams.DiscardUnknown(m)
}

var xxx_messageInfo_ListNVMeParams proto.InternalMessageInfo

// NVMeCtrlr represents a NVMe Controller.
type NVMeCtrlr struct {
	// The model name of the controller.
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// The serial number of the controller.
	Serial               string   `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVMeCtrlr) Reset()         { *m = NVMeCtrlr{} }
func (m *NVMeCtrlr) String() string { return proto.CompactTextString(m) }
func (*NVMeCtrlr) ProtoMessage()    {}
func (*NVMeCtrlr) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{4}
}
func (m *NVMeCtrlr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMeCtrlr.Unmarshal(m, b)
}
func (m *NVMeCtrlr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMeCtrlr.Marshal(b, m, deterministic)
}
func (dst *NVMeCtrlr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMeCtrlr.Merge(dst, src)
}
func (m *NVMeCtrlr) XXX_Size() int {
	return xxx_messageInfo_NVMeCtrlr.Size(m)
}
func (m *NVMeCtrlr) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMeCtrlr.DiscardUnknown(m)
}

var xxx_messageInfo_NVMeCtrlr proto.InternalMessageInfo

func (m *NVMeCtrlr) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *NVMeCtrlr) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

// NVMeNamespace represent NVMe namespaces
// available on controller. Results are
// streamed rather than returned at once.
type NVMeNamespace struct {
	// controller
	Controller *NVMeCtrlr `protobuf:"bytes,1,opt,name=controller,proto3" json:"controller,omitempty"`
	// namespace id
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// device capacity in GBytes
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVMeNamespace) Reset()         { *m = NVMeNamespace{} }
func (m *NVMeNamespace) String() string { return proto.CompactTextString(m) }
func (*NVMeNamespace) ProtoMessage()    {}
func (*NVMeNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_3a0d43eb790a9d64, []int{5}
}
func (m *NVMeNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVMeNamespace.Unmarshal(m, b)
}
func (m *NVMeNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVMeNamespace.Marshal(b, m, deterministic)
}
func (dst *NVMeNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVMeNamespace.Merge(dst, src)
}
func (m *NVMeNamespace) XXX_Size() int {
	return xxx_messageInfo_NVMeNamespace.Size(m)
}
func (m *NVMeNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_NVMeNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_NVMeNamespace proto.InternalMessageInfo

func (m *NVMeNamespace) GetController() *NVMeCtrlr {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *NVMeNamespace) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NVMeNamespace) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func init() {
	proto.RegisterType((*ListFeaturesParams)(nil), "proto.ListFeaturesParams")
	proto.RegisterType((*FeatureName)(nil), "proto.FeatureName")
	proto.RegisterType((*Feature)(nil), "proto.Feature")
	proto.RegisterType((*ListNVMeParams)(nil), "proto.ListNVMeParams")
	proto.RegisterType((*NVMeCtrlr)(nil), "proto.NVMeCtrlr")
	proto.RegisterType((*NVMeNamespace)(nil), "proto.NVMeNamespace")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MgmtControlClient is the client API for MgmtControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MgmtControlClient interface {
	GetFeature(ctx context.Context, in *FeatureName, opts ...grpc.CallOption) (*Feature, error)
	ListFeatures(ctx context.Context, in *ListFeaturesParams, opts ...grpc.CallOption) (MgmtControl_ListFeaturesClient, error)
	ListNVMe(ctx context.Context, in *ListNVMeParams, opts ...grpc.CallOption) (MgmtControl_ListNVMeClient, error)
}

type mgmtControlClient struct {
	cc *grpc.ClientConn
}

func NewMgmtControlClient(cc *grpc.ClientConn) MgmtControlClient {
	return &mgmtControlClient{cc}
}

func (c *mgmtControlClient) GetFeature(ctx context.Context, in *FeatureName, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/proto.MgmtControl/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtControlClient) ListFeatures(ctx context.Context, in *ListFeaturesParams, opts ...grpc.CallOption) (MgmtControl_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MgmtControl_serviceDesc.Streams[0], "/proto.MgmtControl/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &mgmtControlListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MgmtControl_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type mgmtControlListFeaturesClient struct {
	grpc.ClientStream
}

func (x *mgmtControlListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mgmtControlClient) ListNVMe(ctx context.Context, in *ListNVMeParams, opts ...grpc.CallOption) (MgmtControl_ListNVMeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MgmtControl_serviceDesc.Streams[1], "/proto.MgmtControl/ListNVMe", opts...)
	if err != nil {
		return nil, err
	}
	x := &mgmtControlListNVMeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MgmtControl_ListNVMeClient interface {
	Recv() (*NVMeNamespace, error)
	grpc.ClientStream
}

type mgmtControlListNVMeClient struct {
	grpc.ClientStream
}

func (x *mgmtControlListNVMeClient) Recv() (*NVMeNamespace, error) {
	m := new(NVMeNamespace)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MgmtControlServer is the server API for MgmtControl service.
type MgmtControlServer interface {
	GetFeature(context.Context, *FeatureName) (*Feature, error)
	ListFeatures(*ListFeaturesParams, MgmtControl_ListFeaturesServer) error
	ListNVMe(*ListNVMeParams, MgmtControl_ListNVMeServer) error
}

func RegisterMgmtControlServer(s *grpc.Server, srv MgmtControlServer) {
	s.RegisterService(&_MgmtControl_serviceDesc, srv)
}

func _MgmtControl_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtControlServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MgmtControl/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtControlServer).GetFeature(ctx, req.(*FeatureName))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtControl_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFeaturesParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MgmtControlServer).ListFeatures(m, &mgmtControlListFeaturesServer{stream})
}

type MgmtControl_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type mgmtControlListFeaturesServer struct {
	grpc.ServerStream
}

func (x *mgmtControlListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _MgmtControl_ListNVMe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNVMeParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MgmtControlServer).ListNVMe(m, &mgmtControlListNVMeServer{stream})
}

type MgmtControl_ListNVMeServer interface {
	Send(*NVMeNamespace) error
	grpc.ServerStream
}

type mgmtControlListNVMeServer struct {
	grpc.ServerStream
}

func (x *mgmtControlListNVMeServer) Send(m *NVMeNamespace) error {
	return x.ServerStream.SendMsg(m)
}

var _MgmtControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MgmtControl",
	HandlerType: (*MgmtControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _MgmtControl_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _MgmtControl_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListNVMe",
			Handler:       _MgmtControl_ListNVMe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "srv.proto",
}

func init() { proto.RegisterFile("srv.proto", fileDescriptor_srv_3a0d43eb790a9d64) }

var fileDescriptor_srv_3a0d43eb790a9d64 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xf4, 0x30,
	0x14, 0xfc, 0xba, 0x9f, 0x5d, 0xb7, 0xaf, 0x5a, 0x96, 0x47, 0x95, 0xda, 0xd3, 0x9a, 0xd3, 0x9e,
	0x96, 0xa5, 0x9e, 0x44, 0x3c, 0x2d, 0xe8, 0xc5, 0x5d, 0xa4, 0xa0, 0xf7, 0xd8, 0x46, 0x09, 0x34,
	0x4d, 0x49, 0xa2, 0xe0, 0x8f, 0xf3, 0xbf, 0x49, 0xd3, 0x6c, 0x8d, 0xac, 0xa7, 0xf6, 0x4d, 0x66,
	0x26, 0x93, 0x79, 0x10, 0x69, 0xf5, 0xb1, 0xea, 0x94, 0x34, 0x12, 0x43, 0xfb, 0x21, 0x29, 0xe0,
	0x03, 0xd7, 0xe6, 0x8e, 0x51, 0xf3, 0xae, 0x98, 0x7e, 0xa4, 0x8a, 0x0a, 0x4d, 0x2e, 0x21, 0x76,
	0xc8, 0x8e, 0x0a, 0x86, 0x08, 0x47, 0x2d, 0x15, 0x2c, 0x0b, 0x16, 0xc1, 0x32, 0x2a, 0xed, 0x3f,
	0x79, 0x82, 0x63, 0x47, 0xc1, 0x25, 0x84, 0xaf, 0xe3, 0x79, 0x5c, 0xe0, 0x70, 0xc3, 0xca, 0x73,
	0x28, 0x07, 0x02, 0x2e, 0x20, 0xae, 0x99, 0xae, 0x14, 0xef, 0x0c, 0x97, 0x6d, 0x36, 0xb1, 0x7e,
	0x3e, 0x44, 0xe6, 0x90, 0xf4, 0x79, 0x76, 0xcf, 0x5b, 0xe6, 0xb2, 0x5c, 0x43, 0xd4, 0x4f, 0x1b,
	0xa3, 0x1a, 0x85, 0x29, 0x84, 0x42, 0xd6, 0xac, 0x71, 0x51, 0x86, 0x01, 0xcf, 0x61, 0xaa, 0x99,
	0xe2, 0xb4, 0x71, 0x8e, 0x6e, 0x22, 0x02, 0x4e, 0x7b, 0x69, 0x9f, 0x40, 0x77, 0xb4, 0x62, 0xb8,
	0x06, 0xa8, 0x64, 0x6b, 0x94, 0x6c, 0x1a, 0xa6, 0x5c, 0xdc, 0xb9, 0x8b, 0x3b, 0x5e, 0x52, 0x7a,
	0x1c, 0x4c, 0x60, 0xc2, 0x6b, 0x6b, 0x1b, 0x96, 0x13, 0x5e, 0x63, 0x0e, 0xb3, 0x8a, 0x76, 0xb4,
	0xe2, 0xe6, 0x33, 0xfb, 0x6f, 0xd1, 0x71, 0x2e, 0xbe, 0x02, 0x88, 0xb7, 0x6f, 0xc2, 0x6c, 0x06,
	0x39, 0x16, 0x00, 0xf7, 0x6c, 0x5f, 0x2d, 0xfe, 0x51, 0x4b, 0x9e, 0xfc, 0xc6, 0xc8, 0x3f, 0xbc,
	0x85, 0x13, 0x7f, 0x1f, 0x78, 0xe1, 0x18, 0x87, 0x4b, 0x3a, 0x14, 0xaf, 0x03, 0xbc, 0x81, 0xd9,
	0xbe, 0x3e, 0x3c, 0xf3, 0xa4, 0x3f, 0x7d, 0xe6, 0xa9, 0xf7, 0xde, 0xb1, 0x99, 0x5e, 0xfc, 0x32,
	0xb5, 0x07, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xb3, 0x9d, 0x52, 0x26, 0x02, 0x00,
	0x00,
}