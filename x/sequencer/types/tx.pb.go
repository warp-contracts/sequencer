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

type MsgArweaveBlock struct {
	BlockInfo    *ArweaveBlockInfo     `protobuf:"bytes,2,opt,name=block_info,json=blockInfo,proto3" json:"block_info,omitempty"`
	Transactions []*ArweaveTransaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (m *MsgArweaveBlock) Reset()         { *m = MsgArweaveBlock{} }
func (m *MsgArweaveBlock) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveBlock) ProtoMessage()    {}
func (*MsgArweaveBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{2}
}
func (m *MsgArweaveBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveBlock.Merge(m, src)
}
func (m *MsgArweaveBlock) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveBlock proto.InternalMessageInfo

func (m *MsgArweaveBlock) GetBlockInfo() *ArweaveBlockInfo {
	if m != nil {
		return m.BlockInfo
	}
	return nil
}

func (m *MsgArweaveBlock) GetTransactions() []*ArweaveTransaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type MsgArweaveBlockResponse struct {
}

func (m *MsgArweaveBlockResponse) Reset()         { *m = MsgArweaveBlockResponse{} }
func (m *MsgArweaveBlockResponse) String() string { return proto.CompactTextString(m) }
func (*MsgArweaveBlockResponse) ProtoMessage()    {}
func (*MsgArweaveBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca98cc63da9ee56, []int{3}
}
func (m *MsgArweaveBlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgArweaveBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgArweaveBlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgArweaveBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgArweaveBlockResponse.Merge(m, src)
}
func (m *MsgArweaveBlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgArweaveBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgArweaveBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgArweaveBlockResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDataItem)(nil), "sequencer.sequencer.MsgDataItem")
	proto.RegisterType((*MsgDataItemResponse)(nil), "sequencer.sequencer.MsgDataItemResponse")
	proto.RegisterType((*MsgArweaveBlock)(nil), "sequencer.sequencer.MsgArweaveBlock")
	proto.RegisterType((*MsgArweaveBlockResponse)(nil), "sequencer.sequencer.MsgArweaveBlockResponse")
}

func init() { proto.RegisterFile("sequencer/sequencer/tx.proto", fileDescriptor_0ca98cc63da9ee56) }

var fileDescriptor_0ca98cc63da9ee56 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0xaa, 0xda, 0x40,
	0x14, 0xcd, 0x54, 0x28, 0x3a, 0x0a, 0x85, 0xd8, 0x52, 0x0d, 0x25, 0x86, 0xd0, 0xd2, 0x2c, 0x34,
	0x01, 0x4b, 0x77, 0xdd, 0x34, 0xb8, 0x91, 0x22, 0x85, 0x50, 0xba, 0xe8, 0x26, 0x4c, 0xc6, 0x31,
	0x0d, 0x35, 0x33, 0xe9, 0xcc, 0xa4, 0xea, 0x5f, 0xf4, 0x23, 0xfa, 0x23, 0xdd, 0xb9, 0x74, 0x59,
	0xba, 0x90, 0x87, 0xfe, 0xc8, 0x23, 0xc6, 0xc4, 0xf8, 0xf0, 0xf9, 0xde, 0x2a, 0x27, 0xb9, 0xe7,
	0x9e, 0x33, 0xe7, 0xe6, 0x0e, 0x7c, 0x25, 0xc8, 0xcf, 0x94, 0x50, 0x4c, 0xb8, 0x73, 0x42, 0x72,
	0x69, 0x27, 0x9c, 0x49, 0xa6, 0xb6, 0xcb, 0x6f, 0x76, 0x89, 0xb4, 0xe7, 0x21, 0x0b, 0xd9, 0xa1,
	0xee, 0x64, 0x28, 0xa7, 0x6a, 0x5d, 0xcc, 0x44, 0xcc, 0x84, 0x9f, 0x17, 0xf2, 0x97, 0x63, 0xa9,
	0x7f, 0xc9, 0x03, 0xf1, 0x05, 0x41, 0xbf, 0x88, 0x1f, 0xcc, 0x19, 0xfe, 0xe1, 0x47, 0x74, 0x56,
	0x08, 0x0d, 0xae, 0xb1, 0x25, 0x47, 0x54, 0x20, 0x2c, 0x23, 0x46, 0x73, 0xba, 0x99, 0xc0, 0xe6,
	0x44, 0x84, 0x23, 0x24, 0xd1, 0x58, 0x92, 0x58, 0x45, 0xb0, 0x31, 0x45, 0x12, 0xf9, 0x91, 0x24,
	0x71, 0x07, 0x18, 0xc0, 0x6a, 0xb9, 0xa3, 0xf5, 0xb6, 0xa7, 0xfc, 0xdf, 0xf6, 0x3e, 0x84, 0x91,
	0xfc, 0x9e, 0x06, 0x36, 0x66, 0xb1, 0xb3, 0x40, 0x3c, 0x19, 0x60, 0x46, 0x25, 0x47, 0x58, 0x0a,
	0x47, 0xac, 0x72, 0x3f, 0x8e, 0x9d, 0x54, 0x46, 0x73, 0xe1, 0x04, 0x29, 0x9d, 0xce, 0xb9, 0xed,
	0x66, 0x0f, 0x92, 0x09, 0x7b, 0xf5, 0xe9, 0xd1, 0xc2, 0x7c, 0x01, 0xdb, 0x15, 0x47, 0x8f, 0x88,
	0x84, 0x51, 0x41, 0xcc, 0x3f, 0x00, 0x3e, 0x9b, 0x88, 0xf0, 0x63, 0x7e, 0x52, 0x37, 0x8b, 0xa5,
	0x8e, 0x20, 0x3c, 0xe5, 0xeb, 0x3c, 0x31, 0x80, 0xd5, 0x1c, 0xbe, 0xb1, 0x2f, 0x0c, 0xd5, 0xae,
	0xb6, 0x8d, 0xe9, 0x8c, 0x79, 0x8d, 0xa0, 0x80, 0xea, 0x27, 0xd8, 0xaa, 0xe4, 0x16, 0x9d, 0x9a,
	0x51, 0xb3, 0x9a, 0xc3, 0xb7, 0xd7, 0x74, 0xbe, 0x9c, 0xf8, 0xde, 0x59, 0xb3, 0xd9, 0x85, 0x2f,
	0xef, 0x9c, 0xb2, 0x48, 0x30, 0xfc, 0x0b, 0x60, 0x6d, 0x22, 0x42, 0xf5, 0x2b, 0xac, 0x97, 0xf3,
	0x34, 0x2e, 0xba, 0x54, 0xf2, 0x6b, 0xd6, 0x43, 0x8c, 0x42, 0x5f, 0x0d, 0x60, 0xeb, 0x6c, 0x3a,
	0xaf, 0xef, 0xeb, 0xac, 0xb2, 0xb4, 0xfe, 0x63, 0x58, 0x85, 0x87, 0xfb, 0x79, 0xbd, 0xd3, 0xc1,
	0x66, 0xa7, 0x83, 0x9b, 0x9d, 0x0e, 0x7e, 0xef, 0x75, 0x65, 0xb3, 0xd7, 0x95, 0x7f, 0x7b, 0x5d,
	0xf9, 0xf6, 0xfe, 0xca, 0xef, 0x2f, 0xf7, 0x6c, 0x59, 0xbd, 0x05, 0xab, 0x84, 0x88, 0xe0, 0xe9,
	0x61, 0xcd, 0xde, 0xdd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x34, 0x38, 0xd2, 0x75, 0x29, 0x03, 0x00,
	0x00,
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
	ArweaveBlock(ctx context.Context, in *MsgArweaveBlock, opts ...grpc.CallOption) (*MsgArweaveBlockResponse, error)
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

func (c *msgClient) ArweaveBlock(ctx context.Context, in *MsgArweaveBlock, opts ...grpc.CallOption) (*MsgArweaveBlockResponse, error) {
	out := new(MsgArweaveBlockResponse)
	err := c.cc.Invoke(ctx, "/sequencer.sequencer.Msg/ArweaveBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DataItem(context.Context, *MsgDataItem) (*MsgDataItemResponse, error)
	ArweaveBlock(context.Context, *MsgArweaveBlock) (*MsgArweaveBlockResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DataItem(ctx context.Context, req *MsgDataItem) (*MsgDataItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataItem not implemented")
}
func (*UnimplementedMsgServer) ArweaveBlock(ctx context.Context, req *MsgArweaveBlock) (*MsgArweaveBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArweaveBlock not implemented")
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

func _Msg_ArweaveBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgArweaveBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ArweaveBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sequencer.sequencer.Msg/ArweaveBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ArweaveBlock(ctx, req.(*MsgArweaveBlock))
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
			MethodName: "ArweaveBlock",
			Handler:    _Msg_ArweaveBlock_Handler,
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

func (m *MsgArweaveBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.BlockInfo != nil {
		{
			size, err := m.BlockInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgArweaveBlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgArweaveBlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgArweaveBlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgArweaveBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockInfo != nil {
		l = m.BlockInfo.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgArweaveBlockResponse) Size() (n int) {
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
func (m *MsgArweaveBlock) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInfo", wireType)
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
			if m.BlockInfo == nil {
				m.BlockInfo = &ArweaveBlockInfo{}
			}
			if err := m.BlockInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
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
			m.Transactions = append(m.Transactions, &ArweaveTransaction{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgArweaveBlockResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgArweaveBlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgArweaveBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
