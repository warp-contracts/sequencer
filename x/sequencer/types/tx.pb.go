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

type MsgLastArweaveBlock struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Height    uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Hash      []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *MsgLastArweaveBlock) Reset()         { *m = MsgLastArweaveBlock{} }
func (m *MsgLastArweaveBlock) String() string { return proto.CompactTextString(m) }
func (*MsgLastArweaveBlock) ProtoMessage()    {}
func (*MsgLastArweaveBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{2}
}
func (m *MsgLastArweaveBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLastArweaveBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLastArweaveBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLastArweaveBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLastArweaveBlock.Merge(m, src)
}
func (m *MsgLastArweaveBlock) XXX_Size() int {
	return m.Size()
}
func (m *MsgLastArweaveBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLastArweaveBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLastArweaveBlock proto.InternalMessageInfo

func (m *MsgLastArweaveBlock) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgLastArweaveBlock) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgLastArweaveBlock) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgLastArweaveBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type MsgLastArweaveBlockResponse struct {
}

func (m *MsgLastArweaveBlockResponse) Reset()         { *m = MsgLastArweaveBlockResponse{} }
func (m *MsgLastArweaveBlockResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLastArweaveBlockResponse) ProtoMessage()    {}
func (*MsgLastArweaveBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{3}
}
func (m *MsgLastArweaveBlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLastArweaveBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLastArweaveBlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLastArweaveBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLastArweaveBlockResponse.Merge(m, src)
}
func (m *MsgLastArweaveBlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLastArweaveBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLastArweaveBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLastArweaveBlockResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDataItem)(nil), "sequencer.sequencer.MsgDataItem")
	proto.RegisterType((*MsgDataItemResponse)(nil), "sequencer.sequencer.MsgDataItemResponse")
	proto.RegisterType((*MsgLastArweaveBlock)(nil), "sequencer.sequencer.MsgLastArweaveBlock")
	proto.RegisterType((*MsgLastArweaveBlockResponse)(nil), "sequencer.sequencer.MsgLastArweaveBlockResponse")
}

func init() { proto.RegisterFile("sequencer/sequencer/tx.proto", fileDescriptor_0ca98cc63da9ee56) }

var fileDescriptor_0ca98cc63da9ee56 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3d, 0x8f, 0x9b, 0x40,
	0x10, 0x85, 0xd8, 0x72, 0xec, 0x4d, 0x8a, 0x68, 0x9d, 0x44, 0x84, 0x38, 0xd8, 0xa2, 0x72, 0x91,
	0x40, 0x94, 0x28, 0x5d, 0x9a, 0x20, 0x37, 0x91, 0x62, 0x45, 0xa2, 0x48, 0x91, 0x06, 0x2d, 0xeb,
	0x15, 0xa0, 0x00, 0x4b, 0x76, 0x87, 0xb3, 0xfd, 0x2f, 0xee, 0x67, 0xb9, 0x39, 0xc9, 0xe5, 0xe9,
	0x0a, 0xeb, 0x64, 0xff, 0x91, 0x13, 0x60, 0xb0, 0x75, 0xe7, 0xfb, 0xa8, 0x78, 0x33, 0xef, 0xb1,
	0xef, 0xcd, 0xee, 0xa0, 0x81, 0x64, 0xff, 0x73, 0x96, 0x52, 0x26, 0xec, 0x03, 0x82, 0x85, 0x95,
	0x09, 0x0e, 0x1c, 0xf7, 0x9b, 0x9e, 0xd5, 0x20, 0xfd, 0x75, 0xc0, 0x03, 0x5e, 0xf2, 0x76, 0x81,
	0x2a, 0xa9, 0xfe, 0x8e, 0x72, 0x99, 0x70, 0xe9, 0x55, 0x44, 0x55, 0xec, 0xa9, 0x8f, 0xa7, 0x3c,
	0x62, 0x22, 0xc1, 0x23, 0x62, 0xce, 0xc8, 0x19, 0xf3, 0xfc, 0x98, 0xd3, 0x7f, 0x95, 0xda, 0xcc,
	0xd0, 0x8b, 0xa9, 0x0c, 0x26, 0x04, 0xc8, 0x4f, 0x60, 0x09, 0x26, 0xa8, 0x37, 0x23, 0x40, 0xbc,
	0x08, 0x58, 0xa2, 0xa9, 0x23, 0x75, 0xfc, 0xd2, 0x99, 0xac, 0x36, 0x43, 0xe5, 0x6a, 0x33, 0xfc,
	0x1e, 0x44, 0x10, 0xe6, 0xbe, 0x45, 0x79, 0x62, 0xcf, 0x89, 0xc8, 0x3e, 0x51, 0x9e, 0x82, 0x20,
	0x14, 0xa4, 0x2d, 0x97, 0x95, 0x9d, 0xa0, 0x76, 0x0e, 0x51, 0x2c, 0x6d, 0x3f, 0x4f, 0x67, 0xb1,
	0xb0, 0x9c, 0xe2, 0xc3, 0x8a, 0x83, 0xdd, 0xee, 0x6c, 0x6f, 0x61, 0xbe, 0x41, 0xfd, 0x23, 0x47,
	0x97, 0xc9, 0x8c, 0xa7, 0x92, 0x99, 0xcb, 0xb2, 0xfd, 0x8b, 0x48, 0xf8, 0x51, 0xc5, 0x74, 0x8a,
	0x94, 0x58, 0x43, 0xcf, 0xa9, 0x60, 0x04, 0xb8, 0x28, 0xe3, 0xf4, 0xdc, 0xba, 0xc4, 0x6f, 0x51,
	0x27, 0x64, 0x51, 0x10, 0x82, 0xf6, 0x6c, 0xa4, 0x8e, 0xdb, 0xee, 0xbe, 0xc2, 0x03, 0xd4, 0x83,
	0x28, 0x61, 0x12, 0x48, 0x92, 0x69, 0xad, 0x92, 0x3a, 0x34, 0x30, 0x46, 0xed, 0x90, 0xc8, 0x50,
	0x6b, 0x17, 0xb3, 0xb9, 0x25, 0x36, 0x3f, 0xa0, 0xf7, 0x27, 0xac, 0xeb, 0x64, 0x5f, 0x2e, 0x54,
	0xd4, 0x9a, 0xca, 0x00, 0xff, 0x41, 0xdd, 0xe6, 0x9e, 0x46, 0xd6, 0x89, 0xb7, 0xb2, 0x8e, 0xe6,
	0xd2, 0xc7, 0x8f, 0x29, 0xea, 0xf3, 0x71, 0x8a, 0x5e, 0xdd, 0x19, 0xfb, 0xde, 0xbf, 0x6f, 0x2b,
	0xf5, 0xcf, 0x4f, 0x55, 0xd6, 0x7e, 0xce, 0xef, 0xd5, 0xd6, 0x50, 0xd7, 0x5b, 0x43, 0xbd, 0xde,
	0x1a, 0xea, 0xf9, 0xce, 0x50, 0xd6, 0x3b, 0x43, 0xb9, 0xdc, 0x19, 0xca, 0xdf, 0x6f, 0x0f, 0x3c,
	0x71, 0xb3, 0x4a, 0x8b, 0xe3, 0xd5, 0x5d, 0x66, 0x4c, 0xfa, 0x9d, 0x72, 0x95, 0xbe, 0xde, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x63, 0x53, 0x93, 0xde, 0x02, 0x00, 0x00,
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
	LastArweaveBlock(ctx context.Context, in *MsgLastArweaveBlock, opts ...grpc.CallOption) (*MsgLastArweaveBlockResponse, error)
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

func (c *msgClient) LastArweaveBlock(ctx context.Context, in *MsgLastArweaveBlock, opts ...grpc.CallOption) (*MsgLastArweaveBlockResponse, error) {
	out := new(MsgLastArweaveBlockResponse)
	err := c.cc.Invoke(ctx, "/sequencer.sequencer.Msg/LastArweaveBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DataItem(context.Context, *MsgDataItem) (*MsgDataItemResponse, error)
	LastArweaveBlock(context.Context, *MsgLastArweaveBlock) (*MsgLastArweaveBlockResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DataItem(ctx context.Context, req *MsgDataItem) (*MsgDataItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataItem not implemented")
}
func (*UnimplementedMsgServer) LastArweaveBlock(ctx context.Context, req *MsgLastArweaveBlock) (*MsgLastArweaveBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastArweaveBlock not implemented")
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

func _Msg_LastArweaveBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLastArweaveBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LastArweaveBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequencer.sequencer.Msg/LastArweaveBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LastArweaveBlock(ctx, req.(*MsgLastArweaveBlock))
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
			MethodName: "LastArweaveBlock",
			Handler:    _Msg_LastArweaveBlock_Handler,
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

func (m *MsgLastArweaveBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLastArweaveBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLastArweaveBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgLastArweaveBlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLastArweaveBlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLastArweaveBlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgLastArweaveBlock) Size() (n int) {
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

func (m *MsgLastArweaveBlockResponse) Size() (n int) {
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
func (m *MsgLastArweaveBlock) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLastArweaveBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLastArweaveBlock: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgLastArweaveBlockResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgLastArweaveBlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLastArweaveBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
