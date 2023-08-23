// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sequencer/sequencer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_warp_contracts_syncer_src_utils_arweave "github.com/warp-contracts/syncer/src/utils/arweave"
	github_com_warp_contracts_syncer_src_utils_bundlr "github.com/warp-contracts/syncer/src/utils/bundlr"
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

type MsgDataItem struct {
	DataItem github_com_warp_contracts_syncer_src_utils_bundlr.BundleItem `protobuf:"bytes,1,opt,name=data_item,json=dataItem,proto3,customtype=github.com/warp-contracts/syncer/src/utils/bundlr.BundleItem" json:"data_item"`
}

func (m *MsgDataItem) Reset()         { *m = MsgDataItem{} }
func (m *MsgDataItem) String() string { return proto.CompactTextString(m) }
func (*MsgDataItem) ProtoMessage()    {}
func (*MsgDataItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{0}
}
func (m *MsgDataItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDataItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDataItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDataItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDataItem.Merge(m, src)
}
func (m *MsgDataItem) XXX_Size() int {
	return m.Size()
}
func (m *MsgDataItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDataItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDataItem proto.InternalMessageInfo

type MsgDataItemResponse struct {
}

func (m *MsgDataItemResponse) Reset()         { *m = MsgDataItemResponse{} }
func (m *MsgDataItemResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDataItemResponse) ProtoMessage()    {}
func (*MsgDataItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{1}
}
func (m *MsgDataItemResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDataItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDataItemResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDataItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDataItemResponse.Merge(m, src)
}
func (m *MsgDataItemResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDataItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDataItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDataItemResponse proto.InternalMessageInfo

type MsgArweaveBlockInfo struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Height    uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Hash      []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *MsgArweaveBlockInfo) Reset()         { *m = MsgArweaveBlockInfo{} }
func (m *MsgArweaveBlockInfo) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveBlockInfo) ProtoMessage()    {}
func (*MsgArweaveBlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{2}
}
func (m *MsgArweaveBlockInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveBlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveBlockInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveBlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveBlockInfo.Merge(m, src)
}
func (m *MsgArweaveBlockInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveBlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveBlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveBlockInfo proto.InternalMessageInfo

func (m *MsgArweaveBlockInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgArweaveBlockInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgArweaveBlockInfo) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgArweaveBlockInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type MsgArweaveBlockInfoResponse struct {
}

func (m *MsgArweaveBlockInfoResponse) Reset()         { *m = MsgArweaveBlockInfoResponse{} }
func (m *MsgArweaveBlockInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveBlockInfoResponse) ProtoMessage()    {}
func (*MsgArweaveBlockInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{3}
}
func (m *MsgArweaveBlockInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveBlockInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveBlockInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveBlockInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveBlockInfoResponse.Merge(m, src)
}
func (m *MsgArweaveBlockInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveBlockInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveBlockInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveBlockInfoResponse proto.InternalMessageInfo

type MsgArweaveTransaction struct {
	Transaction github_com_warp_contracts_syncer_src_utils_arweave.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3,customtype=github.com/warp-contracts/syncer/src/utils/arweave.Transaction" json:"transaction"`
}

func (m *MsgArweaveTransaction) Reset()         { *m = MsgArweaveTransaction{} }
func (m *MsgArweaveTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveTransaction) ProtoMessage()    {}
func (*MsgArweaveTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{4}
}
func (m *MsgArweaveTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveTransaction.Merge(m, src)
}
func (m *MsgArweaveTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveTransaction proto.InternalMessageInfo

type MsgArweaveTransactionResponse struct {
}

func (m *MsgArweaveTransactionResponse) Reset()         { *m = MsgArweaveTransactionResponse{} }
func (m *MsgArweaveTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveTransactionResponse) ProtoMessage()    {}
func (*MsgArweaveTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{5}
}
func (m *MsgArweaveTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveTransactionResponse.Merge(m, src)
}
func (m *MsgArweaveTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveTransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDataItem)(nil), "sequencer.sequencer.MsgDataItem")
	proto.RegisterType((*MsgDataItemResponse)(nil), "sequencer.sequencer.MsgDataItemResponse")
	proto.RegisterType((*MsgArweaveBlockInfo)(nil), "sequencer.sequencer.MsgArweaveBlockInfo")
	proto.RegisterType((*MsgArweaveBlockInfoResponse)(nil), "sequencer.sequencer.MsgArweaveBlockInfoResponse")
	proto.RegisterType((*MsgArweaveTransaction)(nil), "sequencer.sequencer.MsgArweaveTransaction")
	proto.RegisterType((*MsgArweaveTransactionResponse)(nil), "sequencer.sequencer.MsgArweaveTransactionResponse")
}

func init() { proto.RegisterFile("sequencer/sequencer/tx.proto", fileDescriptor_0ca98cc63da9ee56) }

var fileDescriptor_0ca98cc63da9ee56 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xdb, 0xa8, 0x34, 0x53, 0x0e, 0x68, 0x4b, 0x91, 0x31, 0xad, 0x13, 0xf9, 0x14, 0x21,
	0xb0, 0x51, 0x11, 0x37, 0x84, 0x44, 0x54, 0x21, 0xf5, 0x50, 0x21, 0x59, 0x88, 0x03, 0x97, 0x68,
	0xb3, 0x59, 0x6c, 0x8b, 0x78, 0xd7, 0xec, 0x8e, 0x69, 0x73, 0xe4, 0x1f, 0xf0, 0x9f, 0xb8, 0xf4,
	0xd8, 0x23, 0xe2, 0x50, 0xa1, 0xe4, 0x8f, 0x20, 0x7f, 0x5b, 0x60, 0x68, 0x7a, 0xf2, 0x7c, 0x3c,
	0xcf, 0x7b, 0x7e, 0xe3, 0x81, 0x43, 0xcd, 0x3f, 0xa7, 0x5c, 0x30, 0xae, 0xbc, 0x26, 0xc2, 0x0b,
	0x37, 0x51, 0x12, 0x25, 0xd9, 0xaf, 0x6b, 0x6e, 0x1d, 0x59, 0xf7, 0x03, 0x19, 0xc8, 0xbc, 0xef,
	0x65, 0x51, 0x01, 0xb5, 0x1e, 0x32, 0xa9, 0x63, 0xa9, 0xa7, 0x45, 0xa3, 0x48, 0xca, 0xd6, 0x93,
	0x2e, 0x8e, 0x05, 0xd5, 0x38, 0xa5, 0xea, 0x9c, 0xd3, 0x2f, 0x7c, 0x3a, 0x5b, 0x48, 0xf6, 0xa9,
	0x40, 0x3b, 0x09, 0xec, 0x9d, 0xe9, 0xe0, 0x84, 0x22, 0x3d, 0x45, 0x1e, 0x13, 0x0a, 0x83, 0x39,
	0x45, 0x3a, 0x8d, 0x90, 0xc7, 0xa6, 0x31, 0x32, 0xc6, 0x77, 0x27, 0x27, 0x97, 0xd7, 0xc3, 0xde,
	0xcf, 0xeb, 0xe1, 0xcb, 0x20, 0xc2, 0x30, 0x9d, 0xb9, 0x4c, 0xc6, 0xde, 0x39, 0x55, 0xc9, 0x53,
	0x26, 0x05, 0x2a, 0xca, 0x50, 0x7b, 0x7a, 0x59, 0xd0, 0x29, 0xe6, 0xa5, 0x18, 0x2d, 0xb4, 0x37,
	0x4b, 0xc5, 0x7c, 0xa1, 0xdc, 0x49, 0xf6, 0xe0, 0xd9, 0x60, 0x7f, 0x77, 0x5e, 0x52, 0x38, 0x07,
	0xb0, 0xdf, 0x62, 0xf4, 0xb9, 0x4e, 0xa4, 0xd0, 0xdc, 0x59, 0xe6, 0xe5, 0xd7, 0x85, 0xc4, 0x49,
	0xa6, 0xf0, 0x54, 0x7c, 0x94, 0xc4, 0x84, 0x3b, 0x4c, 0x71, 0x8a, 0x52, 0xe5, 0x72, 0x06, 0x7e,
	0x95, 0x92, 0x07, 0xb0, 0x13, 0xf2, 0x28, 0x08, 0xd1, 0xdc, 0x1a, 0x19, 0xe3, 0xbe, 0x5f, 0x66,
	0xe4, 0x10, 0x06, 0x18, 0xc5, 0x5c, 0x23, 0x8d, 0x13, 0x73, 0x3b, 0x6f, 0x35, 0x05, 0x42, 0xa0,
	0x1f, 0x52, 0x1d, 0x9a, 0xfd, 0xec, 0xdb, 0xfc, 0x3c, 0x76, 0x8e, 0xe0, 0x51, 0x07, 0x75, 0xad,
	0xec, 0xab, 0x01, 0x07, 0x4d, 0xff, 0x9d, 0xa2, 0x42, 0x53, 0x86, 0x91, 0x14, 0x24, 0x84, 0x3d,
	0x6c, 0xd2, 0xd2, 0xaf, 0x37, 0xa5, 0x5f, 0xaf, 0x6e, 0xe1, 0x57, 0xb9, 0x19, 0xb7, 0x35, 0xdc,
	0x6f, 0x8f, 0x76, 0x86, 0x70, 0xd4, 0x29, 0xa1, 0x12, 0x79, 0xfc, 0x7d, 0x0b, 0xb6, 0xcf, 0x74,
	0x40, 0xde, 0xc3, 0x6e, 0xbd, 0xcc, 0x91, 0xdb, 0xf1, 0x43, 0xb9, 0x2d, 0xf3, 0xad, 0xf1, 0x4d,
	0x88, 0x6a, 0x3e, 0x11, 0x70, 0xef, 0xaf, 0xdd, 0xfc, 0xf3, 0xed, 0x3f, 0x91, 0xd6, 0xb3, 0x4d,
	0x91, 0x35, 0x1f, 0x02, 0xe9, 0x30, 0xfc, 0xf1, 0x0d, 0x73, 0x5a, 0x58, 0xeb, 0x78, 0x73, 0x6c,
	0xc5, 0x3a, 0x79, 0x7b, 0xb9, 0xb2, 0x8d, 0xab, 0x95, 0x6d, 0xfc, 0x5a, 0xd9, 0xc6, 0xb7, 0xb5,
	0xdd, 0xbb, 0x5a, 0xdb, 0xbd, 0x1f, 0x6b, 0xbb, 0xf7, 0xe1, 0xc5, 0x7f, 0xb6, 0x59, 0x5f, 0xd9,
	0x45, 0xfb, 0xaa, 0x97, 0x09, 0xd7, 0xb3, 0x9d, 0xfc, 0xca, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xbd, 0x8e, 0x71, 0x63, 0xf9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DataItem(ctx context.Context, in *MsgDataItem, opts ...grpc.CallOption) (*MsgDataItemResponse, error)
	ArweaveBlockInfo(ctx context.Context, in *MsgArweaveBlockInfo, opts ...grpc.CallOption) (*MsgArweaveBlockInfoResponse, error)
	ArweaveTransaction(ctx context.Context, in *MsgArweaveTransaction, opts ...grpc.CallOption) (*MsgArweaveTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DataItem(ctx context.Context, in *MsgDataItem, opts ...grpc.CallOption) (*MsgDataItemResponse, error) {
	out := new(MsgDataItemResponse)
	err := c.cc.Invoke(ctx, "/sequencer.sequencer.Msg/DataItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ArweaveBlockInfo(ctx context.Context, in *MsgArweaveBlockInfo, opts ...grpc.CallOption) (*MsgArweaveBlockInfoResponse, error) {
	out := new(MsgArweaveBlockInfoResponse)
	err := c.cc.Invoke(ctx, "/sequencer.sequencer.Msg/ArweaveBlockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ArweaveTransaction(ctx context.Context, in *MsgArweaveTransaction, opts ...grpc.CallOption) (*MsgArweaveTransactionResponse, error) {
	out := new(MsgArweaveTransactionResponse)
	err := c.cc.Invoke(ctx, "/sequencer.sequencer.Msg/ArweaveTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DataItem(context.Context, *MsgDataItem) (*MsgDataItemResponse, error)
	ArweaveBlockInfo(context.Context, *MsgArweaveBlockInfo) (*MsgArweaveBlockInfoResponse, error)
	ArweaveTransaction(context.Context, *MsgArweaveTransaction) (*MsgArweaveTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DataItem(ctx context.Context, req *MsgDataItem) (*MsgDataItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataItem not implemented")
}
func (*UnimplementedMsgServer) ArweaveBlockInfo(ctx context.Context, req *MsgArweaveBlockInfo) (*MsgArweaveBlockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArweaveBlockInfo not implemented")
}
func (*UnimplementedMsgServer) ArweaveTransaction(ctx context.Context, req *MsgArweaveTransaction) (*MsgArweaveTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArweaveTransaction not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DataItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDataItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DataItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequencer.sequencer.Msg/DataItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DataItem(ctx, req.(*MsgDataItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ArweaveBlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArweaveBlockInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ArweaveBlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequencer.sequencer.Msg/ArweaveBlockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ArweaveBlockInfo(ctx, req.(*MsgArweaveBlockInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ArweaveTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArweaveTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ArweaveTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequencer.sequencer.Msg/ArweaveTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ArweaveTransaction(ctx, req.(*MsgArweaveTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sequencer.sequencer.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataItem",
			Handler:    _Msg_DataItem_Handler,
		},
		{
			MethodName: "ArweaveBlockInfo",
			Handler:    _Msg_ArweaveBlockInfo_Handler,
		},
		{
			MethodName: "ArweaveTransaction",
			Handler:    _Msg_ArweaveTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencer/sequencer/tx.proto",
}

func (m *MsgDataItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDataItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDataItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.DataItem.Size()
		i -= size
		if _, err := m.DataItem.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgDataItemResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDataItemResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDataItemResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgArweaveBlockInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveBlockInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveBlockInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgArweaveBlockInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveBlockInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveBlockInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgArweaveTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Transaction.Size()
		i -= size
		if _, err := m.Transaction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgArweaveTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgDataItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DataItem.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgDataItemResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgArweaveBlockInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTx(uint64(m.Timestamp))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgArweaveBlockInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgArweaveTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Transaction.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgArweaveTransactionResponse) Size() (n int) {
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
func (m *MsgDataItem) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDataItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDataItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataItem", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DataItem.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgDataItemResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDataItemResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDataItemResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgArweaveBlockInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveBlockInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveBlockInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
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
func (m *MsgArweaveBlockInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveBlockInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveBlockInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgArweaveTransaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Transaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgArweaveTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
