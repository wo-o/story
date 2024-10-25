// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/x/evmstaking/types/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgAddWithdrawal represents a message to add a withdrawal request to the withdrawal queue.
type MsgAddWithdrawal struct {
	Authority  string      `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Withdrawal *Withdrawal `protobuf:"bytes,2,opt,name=withdrawal,proto3" json:"withdrawal,omitempty" yaml:"withdrawal"`
}

func (m *MsgAddWithdrawal) Reset()         { *m = MsgAddWithdrawal{} }
func (m *MsgAddWithdrawal) String() string { return proto.CompactTextString(m) }
func (*MsgAddWithdrawal) ProtoMessage()    {}
func (*MsgAddWithdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9acb5c9b0992bb37, []int{0}
}
func (m *MsgAddWithdrawal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddWithdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddWithdrawal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddWithdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddWithdrawal.Merge(m, src)
}
func (m *MsgAddWithdrawal) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddWithdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddWithdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddWithdrawal proto.InternalMessageInfo

func (m *MsgAddWithdrawal) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgAddWithdrawal) GetWithdrawal() *Withdrawal {
	if m != nil {
		return m.Withdrawal
	}
	return nil
}

// MsgAddWithdrawalResponse defines the Msg/AddWithdrawal response type.
type MsgAddWithdrawalResponse struct {
	RequestIndex                uint64 `protobuf:"varint,1,opt,name=request_index,json=requestIndex,proto3" json:"request_index,omitempty" yaml:"request_index"`
	RequestIdDelegatorValidator uint64 `protobuf:"varint,2,opt,name=request_id_delegator_validator,json=requestIdDelegatorValidator,proto3" json:"request_id_delegator_validator,omitempty" yaml:"request_id_delegator_validator"`
}

func (m *MsgAddWithdrawalResponse) Reset()         { *m = MsgAddWithdrawalResponse{} }
func (m *MsgAddWithdrawalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddWithdrawalResponse) ProtoMessage()    {}
func (*MsgAddWithdrawalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9acb5c9b0992bb37, []int{1}
}
func (m *MsgAddWithdrawalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddWithdrawalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddWithdrawalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddWithdrawalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddWithdrawalResponse.Merge(m, src)
}
func (m *MsgAddWithdrawalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddWithdrawalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddWithdrawalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddWithdrawalResponse proto.InternalMessageInfo

func (m *MsgAddWithdrawalResponse) GetRequestIndex() uint64 {
	if m != nil {
		return m.RequestIndex
	}
	return 0
}

func (m *MsgAddWithdrawalResponse) GetRequestIdDelegatorValidator() uint64 {
	if m != nil {
		return m.RequestIdDelegatorValidator
	}
	return 0
}

// MsgRemoveWithdrawal represents a message to remove a withdrawal request from the withdrawal queue.
type MsgRemoveWithdrawal struct {
	Authority                   string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	Delegator                   string `protobuf:"bytes,2,opt,name=delegator,proto3" json:"delegator,omitempty" yaml:"delegator"`
	Validator                   string `protobuf:"bytes,3,opt,name=validator,proto3" json:"validator,omitempty" yaml:"validator"`
	RequestIdDelegatorValidator uint64 `protobuf:"varint,4,opt,name=request_id_delegator_validator,json=requestIdDelegatorValidator,proto3" json:"request_id_delegator_validator,omitempty" yaml:"request_id_delegator_validator"`
}

func (m *MsgRemoveWithdrawal) Reset()         { *m = MsgRemoveWithdrawal{} }
func (m *MsgRemoveWithdrawal) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveWithdrawal) ProtoMessage()    {}
func (*MsgRemoveWithdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9acb5c9b0992bb37, []int{2}
}
func (m *MsgRemoveWithdrawal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveWithdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveWithdrawal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveWithdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveWithdrawal.Merge(m, src)
}
func (m *MsgRemoveWithdrawal) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveWithdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveWithdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveWithdrawal proto.InternalMessageInfo

func (m *MsgRemoveWithdrawal) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgRemoveWithdrawal) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *MsgRemoveWithdrawal) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *MsgRemoveWithdrawal) GetRequestIdDelegatorValidator() uint64 {
	if m != nil {
		return m.RequestIdDelegatorValidator
	}
	return 0
}

// MsgRemoveWithdrawalResponse defines the Msg/RemoveWithdrawal response type.
type MsgRemoveWithdrawalResponse struct {
}

func (m *MsgRemoveWithdrawalResponse) Reset()         { *m = MsgRemoveWithdrawalResponse{} }
func (m *MsgRemoveWithdrawalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveWithdrawalResponse) ProtoMessage()    {}
func (*MsgRemoveWithdrawalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9acb5c9b0992bb37, []int{3}
}
func (m *MsgRemoveWithdrawalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveWithdrawalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveWithdrawalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveWithdrawalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveWithdrawalResponse.Merge(m, src)
}
func (m *MsgRemoveWithdrawalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveWithdrawalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveWithdrawalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveWithdrawalResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddWithdrawal)(nil), "client.x.evmstaking.types.MsgAddWithdrawal")
	proto.RegisterType((*MsgAddWithdrawalResponse)(nil), "client.x.evmstaking.types.MsgAddWithdrawalResponse")
	proto.RegisterType((*MsgRemoveWithdrawal)(nil), "client.x.evmstaking.types.MsgRemoveWithdrawal")
	proto.RegisterType((*MsgRemoveWithdrawalResponse)(nil), "client.x.evmstaking.types.MsgRemoveWithdrawalResponse")
}

func init() {
	proto.RegisterFile("client/x/evmstaking/types/tx.proto", fileDescriptor_9acb5c9b0992bb37)
}

var fileDescriptor_9acb5c9b0992bb37 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0xaf, 0xd0, 0x4f, 0x2d, 0xcb, 0x2d, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x84, 0xa8, 0xd1, 0xab, 0xd0, 0x43, 0xa8, 0xd1, 0x03, 0xab, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0xab, 0xd2, 0x07, 0xb1, 0x20, 0x1a, 0xa4, 0xc4, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5,
	0x73, 0x8b, 0xd3, 0xf5, 0xcb, 0x0c, 0x41, 0x14, 0x54, 0x42, 0x0b, 0xb7, 0x6d, 0x48, 0x46, 0x83,
	0xd5, 0x2a, 0xcd, 0x63, 0xe4, 0x12, 0xf0, 0x2d, 0x4e, 0x77, 0x4c, 0x49, 0x09, 0xcf, 0x2c, 0xc9,
	0x48, 0x29, 0x4a, 0x2c, 0x4f, 0xcc, 0x11, 0x92, 0xe1, 0xe2, 0x4c, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f,
	0xca, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0xc5, 0x70, 0x71,
	0x95, 0xc3, 0xd5, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xea, 0xe1, 0x74, 0xbd, 0x1e,
	0xc2, 0x60, 0x27, 0xd1, 0x4f, 0xf7, 0xe4, 0x05, 0x2b, 0x13, 0x73, 0x73, 0xac, 0x94, 0x10, 0x46,
	0x28, 0x05, 0x21, 0x99, 0x67, 0xc5, 0xd7, 0xf4, 0x7c, 0x83, 0x16, 0xc2, 0x36, 0xa5, 0x93, 0x8c,
	0x5c, 0x12, 0xe8, 0x0e, 0x0c, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xb2, 0xe5, 0xe2,
	0x2d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x00, 0x3b, 0x96,
	0xc5, 0x49, 0xe2, 0xd3, 0x3d, 0x79, 0x11, 0x88, 0x35, 0x28, 0xd2, 0x4a, 0x41, 0x3c, 0x50, 0xbe,
	0x27, 0x88, 0x2b, 0x94, 0xc7, 0x25, 0x07, 0x97, 0x4f, 0x89, 0x4f, 0x49, 0xcd, 0x49, 0x4d, 0x4f,
	0x2c, 0xc9, 0x2f, 0x8a, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x01, 0xb1, 0xc0, 0xbe, 0x63, 0x71, 0xd2,
	0xfc, 0x74, 0x4f, 0x5e, 0x15, 0xcd, 0x3c, 0xac, 0xea, 0x95, 0x82, 0xa4, 0x61, 0x16, 0xa4, 0xb8,
	0xc0, 0xa4, 0xc3, 0xe0, 0xb2, 0xb3, 0x99, 0xb8, 0x84, 0x7d, 0x8b, 0xd3, 0x83, 0x52, 0x73, 0xf3,
	0xcb, 0x52, 0x89, 0x0e, 0x6f, 0x23, 0x2e, 0x4e, 0xb8, 0x55, 0x60, 0x07, 0x71, 0x3a, 0x89, 0x7c,
	0xba, 0x27, 0x2f, 0x00, 0x71, 0x10, 0x5c, 0x4a, 0x29, 0x08, 0xa1, 0x0c, 0xa4, 0x07, 0xe1, 0x09,
	0x66, 0x74, 0x3d, 0x48, 0xee, 0x45, 0x28, 0x23, 0x22, 0x34, 0x58, 0xa8, 0x19, 0x1a, 0x18, 0x31,
	0x2d, 0xcb, 0x25, 0x8d, 0x25, 0x70, 0x60, 0x71, 0x6d, 0xd4, 0xc4, 0xc4, 0xc5, 0xe5, 0x5b, 0x9c,
	0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x54, 0xc8, 0xc5, 0x8b, 0x9a, 0x68, 0xb5, 0xf1,
	0x24, 0x41, 0xf4, 0x04, 0x24, 0x65, 0x4c, 0x82, 0x62, 0x78, 0x6a, 0xab, 0xe2, 0x12, 0xc0, 0x88,
	0x3a, 0x3d, 0xfc, 0x06, 0xa1, 0xab, 0x97, 0x32, 0x23, 0x4d, 0x3d, 0xcc, 0x6e, 0x29, 0xd6, 0x86,
	0xe7, 0x1b, 0xb4, 0x18, 0x9d, 0x8c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x12, 0x67, 0xa6, 0x4f, 0x62, 0x03, 0x67, 0x75, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xef, 0xcf, 0x54, 0x76, 0x86, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// AddWithdrawal defines a method to add a withdrawal to the withdrawal queue.
	AddWithdrawal(ctx context.Context, in *MsgAddWithdrawal, opts ...grpc.CallOption) (*MsgAddWithdrawalResponse, error)
	// RemoveWithdrawal defines a method to remove a withdrawal request from the withdrawal queue.
	RemoveWithdrawal(ctx context.Context, in *MsgRemoveWithdrawal, opts ...grpc.CallOption) (*MsgRemoveWithdrawalResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) AddWithdrawal(ctx context.Context, in *MsgAddWithdrawal, opts ...grpc.CallOption) (*MsgAddWithdrawalResponse, error) {
	out := new(MsgAddWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/client.x.evmstaking.types.MsgService/AddWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RemoveWithdrawal(ctx context.Context, in *MsgRemoveWithdrawal, opts ...grpc.CallOption) (*MsgRemoveWithdrawalResponse, error) {
	out := new(MsgRemoveWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/client.x.evmstaking.types.MsgService/RemoveWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// AddWithdrawal defines a method to add a withdrawal to the withdrawal queue.
	AddWithdrawal(context.Context, *MsgAddWithdrawal) (*MsgAddWithdrawalResponse, error)
	// RemoveWithdrawal defines a method to remove a withdrawal request from the withdrawal queue.
	RemoveWithdrawal(context.Context, *MsgRemoveWithdrawal) (*MsgRemoveWithdrawalResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) AddWithdrawal(ctx context.Context, req *MsgAddWithdrawal) (*MsgAddWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWithdrawal not implemented")
}
func (*UnimplementedMsgServiceServer) RemoveWithdrawal(ctx context.Context, req *MsgRemoveWithdrawal) (*MsgRemoveWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWithdrawal not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_AddWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddWithdrawal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).AddWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.x.evmstaking.types.MsgService/AddWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).AddWithdrawal(ctx, req.(*MsgAddWithdrawal))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RemoveWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveWithdrawal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RemoveWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.x.evmstaking.types.MsgService/RemoveWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RemoveWithdrawal(ctx, req.(*MsgRemoveWithdrawal))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.x.evmstaking.types.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddWithdrawal",
			Handler:    _MsgService_AddWithdrawal_Handler,
		},
		{
			MethodName: "RemoveWithdrawal",
			Handler:    _MsgService_RemoveWithdrawal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/x/evmstaking/types/tx.proto",
}

func (m *MsgAddWithdrawal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddWithdrawal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddWithdrawal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Withdrawal != nil {
		{
			size, err := m.Withdrawal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddWithdrawalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddWithdrawalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddWithdrawalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestIdDelegatorValidator != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestIdDelegatorValidator))
		i--
		dAtA[i] = 0x10
	}
	if m.RequestIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveWithdrawal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveWithdrawal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveWithdrawal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestIdDelegatorValidator != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestIdDelegatorValidator))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveWithdrawalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveWithdrawalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveWithdrawalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddWithdrawal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Withdrawal != nil {
		l = m.Withdrawal.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddWithdrawalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestIndex != 0 {
		n += 1 + sovTx(uint64(m.RequestIndex))
	}
	if m.RequestIdDelegatorValidator != 0 {
		n += 1 + sovTx(uint64(m.RequestIdDelegatorValidator))
	}
	return n
}

func (m *MsgRemoveWithdrawal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.RequestIdDelegatorValidator != 0 {
		n += 1 + sovTx(uint64(m.RequestIdDelegatorValidator))
	}
	return n
}

func (m *MsgRemoveWithdrawalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddWithdrawal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddWithdrawal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddWithdrawal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Withdrawal == nil {
				m.Withdrawal = &Withdrawal{}
			}
			if err := m.Withdrawal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddWithdrawalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddWithdrawalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddWithdrawalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestIndex", wireType)
			}
			m.RequestIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestIdDelegatorValidator", wireType)
			}
			m.RequestIdDelegatorValidator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestIdDelegatorValidator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveWithdrawal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveWithdrawal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveWithdrawal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestIdDelegatorValidator", wireType)
			}
			m.RequestIdDelegatorValidator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestIdDelegatorValidator |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveWithdrawalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveWithdrawalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveWithdrawalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)