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

// message with L2 interaction in the form of a data item
type MsgDataItem struct {
	DataItem    github_com_warp_contracts_syncer_src_utils_bundlr.BundleItem `protobuf:"bytes,1,opt,name=data_item,json=dataItem,proto3,customtype=github.com/warp-contracts/syncer/src/utils/bundlr.BundleItem" json:"data_item"`
	SortKey     string                                                       `protobuf:"bytes,2,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
	LastSortKey string                                                       `protobuf:"bytes,3,opt,name=last_sort_key,json=lastSortKey,proto3" json:"last_sort_key,omitempty"`
	VrfData     *VrfData                                                     `protobuf:"bytes,4,opt,name=vrf_data,json=vrfData,proto3" json:"vrf_data,omitempty"`
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

func (m *MsgDataItem) GetSortKey() string {
	if m != nil {
		return m.SortKey
	}
	return ""
}

func (m *MsgDataItem) GetLastSortKey() string {
	if m != nil {
		return m.LastSortKey
	}
	return ""
}

func (m *MsgDataItem) GetVrfData() *VrfData {
	if m != nil {
		return m.VrfData
	}
	return nil
}

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

// message with an Arweave block
type MsgArweaveBlock struct {
	BlockInfo    *ArweaveBlockInfo                    `protobuf:"bytes,1,opt,name=block_info,json=blockInfo,proto3" json:"block_info,omitempty"`
	Transactions []*ArweaveTransactionWithLastSortKey `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
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

func (m *MsgArweaveBlock) GetTransactions() []*ArweaveTransactionWithLastSortKey {
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
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x8f, 0xd2, 0x40,
	0x18, 0x65, 0x16, 0xe3, 0xc2, 0x14, 0x63, 0xd2, 0xd5, 0x58, 0xc8, 0xda, 0x6d, 0x1a, 0x4d, 0x7a,
	0x58, 0xdb, 0x04, 0xa3, 0x5e, 0x3c, 0x28, 0xe1, 0xb2, 0x51, 0x62, 0x52, 0xcd, 0x9a, 0xec, 0x65,
	0x32, 0x2d, 0xd3, 0xd2, 0x2c, 0x74, 0x70, 0xe6, 0x03, 0x96, 0x7f, 0xe1, 0x7f, 0xf1, 0x17, 0x78,
	0xdb, 0xe3, 0x1e, 0x8d, 0x07, 0x62, 0xe0, 0xe2, 0xcf, 0x30, 0xa5, 0x94, 0x16, 0x53, 0x71, 0x4f,
	0xfd, 0x66, 0xbe, 0x37, 0xef, 0xcd, 0xfb, 0xde, 0x14, 0x1f, 0x4b, 0xf6, 0x65, 0xc2, 0x62, 0x9f,
	0x09, 0x27, 0xaf, 0xe0, 0xca, 0x1e, 0x0b, 0x0e, 0x5c, 0x3d, 0xda, 0xee, 0xd9, 0xdb, 0xaa, 0xf5,
	0x20, 0xe4, 0x21, 0x5f, 0xf7, 0x9d, 0xa4, 0x4a, 0xa1, 0xad, 0xa6, 0xcf, 0xe5, 0x88, 0x4b, 0x92,
	0x36, 0xd2, 0xc5, 0xa6, 0x75, 0x5a, 0xa6, 0x41, 0xc5, 0x8c, 0xd1, 0x29, 0x23, 0xde, 0x90, 0xfb,
	0x97, 0x24, 0x8a, 0x83, 0x8c, 0xe8, 0xcd, 0x3e, 0x34, 0x08, 0x1a, 0x4b, 0xea, 0x43, 0xc4, 0x63,
	0x32, 0x8b, 0x60, 0x40, 0x86, 0x54, 0x02, 0x91, 0x5c, 0x00, 0xb9, 0x64, 0xf3, 0x0d, 0xc3, 0xe3,
	0x32, 0x86, 0xa9, 0x08, 0xd2, 0xb6, 0xf9, 0x1b, 0x61, 0xa5, 0x27, 0xc3, 0x2e, 0x05, 0x7a, 0x06,
	0x6c, 0xa4, 0x52, 0x5c, 0xef, 0x53, 0xa0, 0x24, 0x02, 0x36, 0xd2, 0x90, 0x81, 0xac, 0x46, 0xa7,
	0x7b, 0xbd, 0x38, 0xa9, 0xfc, 0x5c, 0x9c, 0xbc, 0x0e, 0x23, 0x18, 0x4c, 0x3c, 0xdb, 0xe7, 0x23,
	0x67, 0x46, 0xc5, 0xf8, 0x99, 0xcf, 0x63, 0x10, 0xd4, 0x07, 0xe9, 0xc8, 0x79, 0x2a, 0x20, 0x7c,
	0x67, 0x02, 0xd1, 0x50, 0x3a, 0xde, 0x24, 0xee, 0x0f, 0x85, 0xdd, 0x49, 0x3e, 0x2c, 0x21, 0x76,
	0x6b, 0xfd, 0x4c, 0xa2, 0x89, 0x6b, 0xd9, 0x1d, 0xb5, 0x03, 0x03, 0x59, 0x75, 0xf7, 0x30, 0x59,
	0xbf, 0x63, 0x73, 0xd5, 0xc4, 0xf7, 0x76, 0x3c, 0x68, 0xd5, 0x75, 0x5f, 0x49, 0x36, 0x3f, 0x6e,
	0x30, 0xaf, 0x70, 0x6d, 0x2a, 0x02, 0x92, 0xd0, 0x69, 0x77, 0x0c, 0x64, 0x29, 0xed, 0x63, 0xbb,
	0x24, 0x19, 0xfb, 0x5c, 0x04, 0x89, 0x2b, 0xf7, 0x70, 0x9a, 0x16, 0xe6, 0x43, 0x7c, 0x54, 0x70,
	0xea, 0x32, 0x39, 0xe6, 0xb1, 0x64, 0xe6, 0x37, 0x84, 0xef, 0xf7, 0x64, 0xf8, 0x36, 0x1d, 0x6a,
	0x27, 0x49, 0x40, 0xed, 0x62, 0x9c, 0x47, 0xb1, 0x1e, 0x83, 0xd2, 0x7e, 0x5a, 0xaa, 0x52, 0x3c,
	0x76, 0x16, 0x07, 0xdc, 0xad, 0x7b, 0x59, 0xa9, 0x5e, 0xe0, 0x46, 0x21, 0x22, 0xa9, 0x1d, 0x18,
	0x55, 0x4b, 0x69, 0xbf, 0xdc, 0xc7, 0xf3, 0x29, 0xc7, 0x7f, 0x8e, 0x60, 0xf0, 0x3e, 0xf7, 0xed,
	0xee, 0x70, 0x99, 0x4d, 0xfc, 0xe8, 0xaf, 0x4b, 0x67, 0x86, 0xda, 0xdf, 0x11, 0xae, 0xf6, 0x64,
	0xa8, 0x9e, 0xe3, 0xda, 0x36, 0x56, 0xa3, 0x54, 0xb4, 0x30, 0x8e, 0x96, 0xf5, 0x3f, 0x44, 0xc6,
	0xaf, 0x7a, 0xb8, 0xb1, 0x33, 0xac, 0x27, 0xff, 0x3a, 0x59, 0x44, 0xb5, 0x4e, 0x6f, 0x83, 0xca,
	0x34, 0x3a, 0x1f, 0xae, 0x97, 0x3a, 0xba, 0x59, 0xea, 0xe8, 0xd7, 0x52, 0x47, 0x5f, 0x57, 0x7a,
	0xe5, 0x66, 0xa5, 0x57, 0x7e, 0xac, 0xf4, 0xca, 0xc5, 0x8b, 0x3d, 0xaf, 0x70, 0xfb, 0xbe, 0xaf,
	0x8a, 0xff, 0xef, 0x7c, 0xcc, 0xa4, 0x77, 0x77, 0xfd, 0xdc, 0x9f, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0x8f, 0x86, 0x20, 0xe3, 0x03, 0x00, 0x00,
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
	if m.VrfData != nil {
		{
			size, err := m.VrfData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.LastSortKey) > 0 {
		i -= len(m.LastSortKey)
		copy(dAtA[i:], m.LastSortKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.LastSortKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SortKey) > 0 {
		i -= len(m.SortKey)
		copy(dAtA[i:], m.SortKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SortKey)))
		i--
		dAtA[i] = 0x12
	}
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
			dAtA[i] = 0x12
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
		dAtA[i] = 0xa
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
	l = len(m.SortKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.LastSortKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VrfData != nil {
		l = m.VrfData.Size()
		n += 1 + l + sovTx(uint64(l))
	}
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortKey", wireType)
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
			m.SortKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSortKey", wireType)
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
			m.LastSortKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfData", wireType)
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
			if m.VrfData == nil {
				m.VrfData = &VrfData{}
			}
			if err := m.VrfData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 1:
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
		case 2:
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
			m.Transactions = append(m.Transactions, &ArweaveTransactionWithLastSortKey{})
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
