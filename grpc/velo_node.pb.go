// Code generated by protoc-gen-go. DO NOT EDIT.
// source: velo_node.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VeloTxRequest struct {
	SignedVeloTxXdr      string   `protobuf:"bytes,1,opt,name=signedVeloTxXdr,proto3" json:"signedVeloTxXdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VeloTxRequest) Reset()         { *m = VeloTxRequest{} }
func (m *VeloTxRequest) String() string { return proto.CompactTextString(m) }
func (*VeloTxRequest) ProtoMessage()    {}
func (*VeloTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a824ed0044bf442b, []int{0}
}

func (m *VeloTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VeloTxRequest.Unmarshal(m, b)
}
func (m *VeloTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VeloTxRequest.Marshal(b, m, deterministic)
}
func (m *VeloTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VeloTxRequest.Merge(m, src)
}
func (m *VeloTxRequest) XXX_Size() int {
	return xxx_messageInfo_VeloTxRequest.Size(m)
}
func (m *VeloTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VeloTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VeloTxRequest proto.InternalMessageInfo

func (m *VeloTxRequest) GetSignedVeloTxXdr() string {
	if m != nil {
		return m.SignedVeloTxXdr
	}
	return ""
}

type VeloTxReply struct {
	SignedStellarTxXdr   string   `protobuf:"bytes,1,opt,name=signedStellarTxXdr,proto3" json:"signedStellarTxXdr,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VeloTxReply) Reset()         { *m = VeloTxReply{} }
func (m *VeloTxReply) String() string { return proto.CompactTextString(m) }
func (*VeloTxReply) ProtoMessage()    {}
func (*VeloTxReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a824ed0044bf442b, []int{1}
}

func (m *VeloTxReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VeloTxReply.Unmarshal(m, b)
}
func (m *VeloTxReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VeloTxReply.Marshal(b, m, deterministic)
}
func (m *VeloTxReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VeloTxReply.Merge(m, src)
}
func (m *VeloTxReply) XXX_Size() int {
	return xxx_messageInfo_VeloTxReply.Size(m)
}
func (m *VeloTxReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VeloTxReply.DiscardUnknown(m)
}

var xxx_messageInfo_VeloTxReply proto.InternalMessageInfo

func (m *VeloTxReply) GetSignedStellarTxXdr() string {
	if m != nil {
		return m.SignedStellarTxXdr
	}
	return ""
}

func (m *VeloTxReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetExchangeRateRequest struct {
	AssetCode            string   `protobuf:"bytes,1,opt,name=assetCode,proto3" json:"assetCode,omitempty"`
	Issuer               string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetExchangeRateRequest) Reset()         { *m = GetExchangeRateRequest{} }
func (m *GetExchangeRateRequest) String() string { return proto.CompactTextString(m) }
func (*GetExchangeRateRequest) ProtoMessage()    {}
func (*GetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a824ed0044bf442b, []int{2}
}

func (m *GetExchangeRateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExchangeRateRequest.Unmarshal(m, b)
}
func (m *GetExchangeRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExchangeRateRequest.Marshal(b, m, deterministic)
}
func (m *GetExchangeRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExchangeRateRequest.Merge(m, src)
}
func (m *GetExchangeRateRequest) XXX_Size() int {
	return xxx_messageInfo_GetExchangeRateRequest.Size(m)
}
func (m *GetExchangeRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExchangeRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExchangeRateRequest proto.InternalMessageInfo

func (m *GetExchangeRateRequest) GetAssetCode() string {
	if m != nil {
		return m.AssetCode
	}
	return ""
}

func (m *GetExchangeRateRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

type GetExchangeRateReply struct {
	AssetCode              string   `protobuf:"bytes,1,opt,name=assetCode,proto3" json:"assetCode,omitempty"`
	Issuer                 string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	RedeemableCollateral   string   `protobuf:"bytes,3,opt,name=redeemableCollateral,proto3" json:"redeemableCollateral,omitempty"`
	RedeemablePricePerUnit string   `protobuf:"bytes,4,opt,name=redeemablePricePerUnit,proto3" json:"redeemablePricePerUnit,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetExchangeRateReply) Reset()         { *m = GetExchangeRateReply{} }
func (m *GetExchangeRateReply) String() string { return proto.CompactTextString(m) }
func (*GetExchangeRateReply) ProtoMessage()    {}
func (*GetExchangeRateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a824ed0044bf442b, []int{3}
}

func (m *GetExchangeRateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExchangeRateReply.Unmarshal(m, b)
}
func (m *GetExchangeRateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExchangeRateReply.Marshal(b, m, deterministic)
}
func (m *GetExchangeRateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExchangeRateReply.Merge(m, src)
}
func (m *GetExchangeRateReply) XXX_Size() int {
	return xxx_messageInfo_GetExchangeRateReply.Size(m)
}
func (m *GetExchangeRateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExchangeRateReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetExchangeRateReply proto.InternalMessageInfo

func (m *GetExchangeRateReply) GetAssetCode() string {
	if m != nil {
		return m.AssetCode
	}
	return ""
}

func (m *GetExchangeRateReply) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *GetExchangeRateReply) GetRedeemableCollateral() string {
	if m != nil {
		return m.RedeemableCollateral
	}
	return ""
}

func (m *GetExchangeRateReply) GetRedeemablePricePerUnit() string {
	if m != nil {
		return m.RedeemablePricePerUnit
	}
	return ""
}

func init() {
	proto.RegisterType((*VeloTxRequest)(nil), "grpc.VeloTxRequest")
	proto.RegisterType((*VeloTxReply)(nil), "grpc.VeloTxReply")
	proto.RegisterType((*GetExchangeRateRequest)(nil), "grpc.GetExchangeRateRequest")
	proto.RegisterType((*GetExchangeRateReply)(nil), "grpc.GetExchangeRateReply")
}

func init() { proto.RegisterFile("velo_node.proto", fileDescriptor_a824ed0044bf442b) }

var fileDescriptor_a824ed0044bf442b = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x35, 0x5a, 0xaa, 0xbd, 0x2a, 0xc1, 0x6b, 0x09, 0x21, 0x74, 0x21, 0x59, 0x75, 0x95, 0x45,
	0x05, 0xd1, 0x75, 0x11, 0x57, 0x96, 0x92, 0xfa, 0xda, 0xc9, 0x24, 0xb9, 0xc4, 0xc0, 0x24, 0x13,
	0x67, 0x26, 0xd2, 0x7e, 0x87, 0x3f, 0xe3, 0xe7, 0x49, 0x5e, 0xd4, 0xd6, 0x74, 0xd3, 0xe5, 0x9c,
	0x17, 0x87, 0x73, 0x07, 0xcc, 0x2f, 0xe2, 0xe2, 0x3d, 0x13, 0x11, 0x79, 0xb9, 0x14, 0x5a, 0x60,
	0x2f, 0x96, 0x79, 0xe8, 0xde, 0xc1, 0xf9, 0x0b, 0x71, 0xf1, 0xb4, 0xf4, 0xe9, 0xb3, 0x20, 0xa5,
	0x71, 0x0c, 0xa6, 0x4a, 0xe2, 0x8c, 0xa2, 0x1a, 0x7e, 0x8b, 0xa4, 0x6d, 0x5c, 0x19, 0xe3, 0x81,
	0xbf, 0x0d, 0xbb, 0xaf, 0x70, 0xda, 0x5a, 0x73, 0xbe, 0x42, 0x0f, 0xb0, 0x56, 0x2c, 0x34, 0x71,
	0xce, 0xe4, 0x5f, 0x6f, 0x07, 0x83, 0x36, 0x1c, 0xa7, 0xa4, 0x14, 0x8b, 0xc9, 0x3e, 0xac, 0x44,
	0xed, 0xd3, 0x9d, 0x81, 0xf5, 0x40, 0xfa, 0x7e, 0x19, 0x7e, 0xb0, 0x2c, 0x26, 0x9f, 0x69, 0x6a,
	0xcb, 0x8d, 0x60, 0xc0, 0x94, 0x22, 0x3d, 0x15, 0x11, 0x35, 0xd1, 0x6b, 0x00, 0x2d, 0xe8, 0x27,
	0x4a, 0x15, 0x24, 0x9b, 0xc0, 0xe6, 0xe5, 0xfe, 0x18, 0x30, 0xfc, 0x17, 0x58, 0x56, 0xde, 0x2b,
	0x0e, 0x27, 0x30, 0x94, 0x14, 0x11, 0xa5, 0x2c, 0xe0, 0x34, 0x15, 0x9c, 0x33, 0x4d, 0x92, 0x71,
	0xfb, 0xa8, 0x52, 0x75, 0x72, 0x78, 0x03, 0xd6, 0x1a, 0x9f, 0xcb, 0x24, 0xa4, 0x39, 0xc9, 0xe7,
	0x2c, 0xd1, 0x76, 0xaf, 0x72, 0xed, 0x60, 0x27, 0xdf, 0x06, 0x9c, 0x94, 0x23, 0xcf, 0xca, 0x42,
	0xb7, 0x70, 0xb6, 0x28, 0x82, 0x34, 0xd1, 0xf5, 0xec, 0x78, 0xe9, 0x95, 0x27, 0xf4, 0x36, 0xee,
	0xe7, 0x5c, 0x6c, 0x82, 0x39, 0x5f, 0xb9, 0x07, 0xf8, 0x08, 0xe6, 0xd6, 0x00, 0x38, 0xaa, 0x75,
	0xdd, 0x43, 0x3b, 0xce, 0x0e, 0xb6, 0x8a, 0x0b, 0xfa, 0xd5, 0x0f, 0xba, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0x9e, 0x51, 0xbc, 0x54, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VeloNodeClient is the client API for VeloNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VeloNodeClient interface {
	SubmitVeloTx(ctx context.Context, in *VeloTxRequest, opts ...grpc.CallOption) (*VeloTxReply, error)
	GetExchangeRate(ctx context.Context, in *GetExchangeRateRequest, opts ...grpc.CallOption) (*GetExchangeRateReply, error)
}

type veloNodeClient struct {
	cc *grpc.ClientConn
}

func NewVeloNodeClient(cc *grpc.ClientConn) VeloNodeClient {
	return &veloNodeClient{cc}
}

func (c *veloNodeClient) SubmitVeloTx(ctx context.Context, in *VeloTxRequest, opts ...grpc.CallOption) (*VeloTxReply, error) {
	out := new(VeloTxReply)
	err := c.cc.Invoke(ctx, "/grpc.VeloNode/SubmitVeloTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *veloNodeClient) GetExchangeRate(ctx context.Context, in *GetExchangeRateRequest, opts ...grpc.CallOption) (*GetExchangeRateReply, error) {
	out := new(GetExchangeRateReply)
	err := c.cc.Invoke(ctx, "/grpc.VeloNode/GetExchangeRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VeloNodeServer is the server API for VeloNode service.
type VeloNodeServer interface {
	SubmitVeloTx(context.Context, *VeloTxRequest) (*VeloTxReply, error)
	GetExchangeRate(context.Context, *GetExchangeRateRequest) (*GetExchangeRateReply, error)
}

// UnimplementedVeloNodeServer can be embedded to have forward compatible implementations.
type UnimplementedVeloNodeServer struct {
}

func (*UnimplementedVeloNodeServer) SubmitVeloTx(ctx context.Context, req *VeloTxRequest) (*VeloTxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVeloTx not implemented")
}
func (*UnimplementedVeloNodeServer) GetExchangeRate(ctx context.Context, req *GetExchangeRateRequest) (*GetExchangeRateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRate not implemented")
}

func RegisterVeloNodeServer(s *grpc.Server, srv VeloNodeServer) {
	s.RegisterService(&_VeloNode_serviceDesc, srv)
}

func _VeloNode_SubmitVeloTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VeloTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VeloNodeServer).SubmitVeloTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VeloNode/SubmitVeloTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VeloNodeServer).SubmitVeloTx(ctx, req.(*VeloTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VeloNode_GetExchangeRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExchangeRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VeloNodeServer).GetExchangeRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.VeloNode/GetExchangeRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VeloNodeServer).GetExchangeRate(ctx, req.(*GetExchangeRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VeloNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.VeloNode",
	HandlerType: (*VeloNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitVeloTx",
			Handler:    _VeloNode_SubmitVeloTx_Handler,
		},
		{
			MethodName: "GetExchangeRate",
			Handler:    _VeloNode_GetExchangeRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "velo_node.proto",
}
