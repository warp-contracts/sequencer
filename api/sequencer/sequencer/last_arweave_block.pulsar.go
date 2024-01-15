// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package sequencer

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_LastArweaveBlock                        protoreflect.MessageDescriptor
	fd_LastArweaveBlock_arweave_block          protoreflect.FieldDescriptor
	fd_LastArweaveBlock_sequencer_block_height protoreflect.FieldDescriptor
)

func init() {
	file_sequencer_sequencer_last_arweave_block_proto_init()
	md_LastArweaveBlock = File_sequencer_sequencer_last_arweave_block_proto.Messages().ByName("LastArweaveBlock")
	fd_LastArweaveBlock_arweave_block = md_LastArweaveBlock.Fields().ByName("arweave_block")
	fd_LastArweaveBlock_sequencer_block_height = md_LastArweaveBlock.Fields().ByName("sequencer_block_height")
}

var _ protoreflect.Message = (*fastReflection_LastArweaveBlock)(nil)

type fastReflection_LastArweaveBlock LastArweaveBlock

func (x *LastArweaveBlock) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LastArweaveBlock)(x)
}

func (x *LastArweaveBlock) slowProtoReflect() protoreflect.Message {
	mi := &file_sequencer_sequencer_last_arweave_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LastArweaveBlock_messageType fastReflection_LastArweaveBlock_messageType
var _ protoreflect.MessageType = fastReflection_LastArweaveBlock_messageType{}

type fastReflection_LastArweaveBlock_messageType struct{}

func (x fastReflection_LastArweaveBlock_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LastArweaveBlock)(nil)
}
func (x fastReflection_LastArweaveBlock_messageType) New() protoreflect.Message {
	return new(fastReflection_LastArweaveBlock)
}
func (x fastReflection_LastArweaveBlock_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LastArweaveBlock
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LastArweaveBlock) Descriptor() protoreflect.MessageDescriptor {
	return md_LastArweaveBlock
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LastArweaveBlock) Type() protoreflect.MessageType {
	return _fastReflection_LastArweaveBlock_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LastArweaveBlock) New() protoreflect.Message {
	return new(fastReflection_LastArweaveBlock)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LastArweaveBlock) Interface() protoreflect.ProtoMessage {
	return (*LastArweaveBlock)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LastArweaveBlock) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ArweaveBlock != nil {
		value := protoreflect.ValueOfMessage(x.ArweaveBlock.ProtoReflect())
		if !f(fd_LastArweaveBlock_arweave_block, value) {
			return
		}
	}
	if x.SequencerBlockHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.SequencerBlockHeight)
		if !f(fd_LastArweaveBlock_sequencer_block_height, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_LastArweaveBlock) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		return x.ArweaveBlock != nil
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		return x.SequencerBlockHeight != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastArweaveBlock) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		x.ArweaveBlock = nil
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		x.SequencerBlockHeight = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LastArweaveBlock) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		value := x.ArweaveBlock
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		value := x.SequencerBlockHeight
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastArweaveBlock) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		x.ArweaveBlock = value.Message().Interface().(*ArweaveBlockInfo)
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		x.SequencerBlockHeight = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastArweaveBlock) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		if x.ArweaveBlock == nil {
			x.ArweaveBlock = new(ArweaveBlockInfo)
		}
		return protoreflect.ValueOfMessage(x.ArweaveBlock.ProtoReflect())
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		panic(fmt.Errorf("field sequencer_block_height of message sequencer.sequencer.LastArweaveBlock is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LastArweaveBlock) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "sequencer.sequencer.LastArweaveBlock.arweave_block":
		m := new(ArweaveBlockInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "sequencer.sequencer.LastArweaveBlock.sequencer_block_height":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: sequencer.sequencer.LastArweaveBlock"))
		}
		panic(fmt.Errorf("message sequencer.sequencer.LastArweaveBlock does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LastArweaveBlock) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in sequencer.sequencer.LastArweaveBlock", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LastArweaveBlock) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LastArweaveBlock) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_LastArweaveBlock) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LastArweaveBlock) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LastArweaveBlock)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ArweaveBlock != nil {
			l = options.Size(x.ArweaveBlock)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.SequencerBlockHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.SequencerBlockHeight))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*LastArweaveBlock)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.SequencerBlockHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SequencerBlockHeight))
			i--
			dAtA[i] = 0x10
		}
		if x.ArweaveBlock != nil {
			encoded, err := options.Marshal(x.ArweaveBlock)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*LastArweaveBlock)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LastArweaveBlock: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LastArweaveBlock: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ArweaveBlock", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ArweaveBlock == nil {
					x.ArweaveBlock = &ArweaveBlockInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ArweaveBlock); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SequencerBlockHeight", wireType)
				}
				x.SequencerBlockHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SequencerBlockHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sequencer/sequencer/last_arweave_block.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information about the latest Arweave block added to the blockchain
type LastArweaveBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArweaveBlock         *ArweaveBlockInfo `protobuf:"bytes,1,opt,name=arweave_block,json=arweaveBlock,proto3" json:"arweave_block,omitempty"`
	SequencerBlockHeight int64             `protobuf:"varint,2,opt,name=sequencer_block_height,json=sequencerBlockHeight,proto3" json:"sequencer_block_height,omitempty"`
}

func (x *LastArweaveBlock) Reset() {
	*x = LastArweaveBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sequencer_sequencer_last_arweave_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastArweaveBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastArweaveBlock) ProtoMessage() {}

// Deprecated: Use LastArweaveBlock.ProtoReflect.Descriptor instead.
func (*LastArweaveBlock) Descriptor() ([]byte, []int) {
	return file_sequencer_sequencer_last_arweave_block_proto_rawDescGZIP(), []int{0}
}

func (x *LastArweaveBlock) GetArweaveBlock() *ArweaveBlockInfo {
	if x != nil {
		return x.ArweaveBlock
	}
	return nil
}

func (x *LastArweaveBlock) GetSequencerBlockHeight() int64 {
	if x != nil {
		return x.SequencerBlockHeight
	}
	return 0
}

var File_sequencer_sequencer_last_arweave_block_proto protoreflect.FileDescriptor

var file_sequencer_sequencer_last_arweave_block_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x72, 0x77, 0x65, 0x61,
	0x76, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0x1a, 0x2c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x01, 0x0a, 0x10, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x4a, 0x0a, 0x0d, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x61, 0x72, 0x77, 0x65, 0x61, 0x76, 0x65, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x14, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0xc3, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x72, 0x42, 0x15, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x72, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0xca,
	0x02, 0x13, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x72, 0xe2, 0x02, 0x1f, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x72, 0x5c, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sequencer_sequencer_last_arweave_block_proto_rawDescOnce sync.Once
	file_sequencer_sequencer_last_arweave_block_proto_rawDescData = file_sequencer_sequencer_last_arweave_block_proto_rawDesc
)

func file_sequencer_sequencer_last_arweave_block_proto_rawDescGZIP() []byte {
	file_sequencer_sequencer_last_arweave_block_proto_rawDescOnce.Do(func() {
		file_sequencer_sequencer_last_arweave_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_sequencer_sequencer_last_arweave_block_proto_rawDescData)
	})
	return file_sequencer_sequencer_last_arweave_block_proto_rawDescData
}

var file_sequencer_sequencer_last_arweave_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sequencer_sequencer_last_arweave_block_proto_goTypes = []interface{}{
	(*LastArweaveBlock)(nil), // 0: sequencer.sequencer.LastArweaveBlock
	(*ArweaveBlockInfo)(nil), // 1: sequencer.sequencer.ArweaveBlockInfo
}
var file_sequencer_sequencer_last_arweave_block_proto_depIdxs = []int32{
	1, // 0: sequencer.sequencer.LastArweaveBlock.arweave_block:type_name -> sequencer.sequencer.ArweaveBlockInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sequencer_sequencer_last_arweave_block_proto_init() }
func file_sequencer_sequencer_last_arweave_block_proto_init() {
	if File_sequencer_sequencer_last_arweave_block_proto != nil {
		return
	}
	file_sequencer_sequencer_arweave_block_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sequencer_sequencer_last_arweave_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastArweaveBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sequencer_sequencer_last_arweave_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sequencer_sequencer_last_arweave_block_proto_goTypes,
		DependencyIndexes: file_sequencer_sequencer_last_arweave_block_proto_depIdxs,
		MessageInfos:      file_sequencer_sequencer_last_arweave_block_proto_msgTypes,
	}.Build()
	File_sequencer_sequencer_last_arweave_block_proto = out.File
	file_sequencer_sequencer_last_arweave_block_proto_rawDesc = nil
	file_sequencer_sequencer_last_arweave_block_proto_goTypes = nil
	file_sequencer_sequencer_last_arweave_block_proto_depIdxs = nil
}
