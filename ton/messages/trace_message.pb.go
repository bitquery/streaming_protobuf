// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: ton/trace_message.proto

package ton_messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeader *BlockHeader `protobuf:"bytes,1,opt,name=BlockHeader,proto3" json:"BlockHeader,omitempty"`
	Transaction *Transaction `protobuf:"bytes,2,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	Depth       uint32       `protobuf:"varint,3,opt,name=Depth,proto3" json:"Depth,omitempty"`
	Children    []*Trace     `protobuf:"bytes,4,rep,name=Children,proto3" json:"Children,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	mi := &file_ton_trace_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_ton_trace_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_ton_trace_message_proto_rawDescGZIP(), []int{0}
}

func (x *Trace) GetBlockHeader() *BlockHeader {
	if x != nil {
		return x.BlockHeader
	}
	return nil
}

func (x *Trace) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Trace) GetDepth() uint32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *Trace) GetChildren() []*Trace {
	if x != nil {
		return x.Children
	}
	return nil
}

type TraceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain  *Chain   `protobuf:"bytes,1,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Traces []*Trace `protobuf:"bytes,2,rep,name=Traces,proto3" json:"Traces,omitempty"`
}

func (x *TraceMessage) Reset() {
	*x = TraceMessage{}
	mi := &file_ton_trace_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceMessage) ProtoMessage() {}

func (x *TraceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ton_trace_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceMessage.ProtoReflect.Descriptor instead.
func (*TraceMessage) Descriptor() ([]byte, []int) {
	return file_ton_trace_message_proto_rawDescGZIP(), []int{1}
}

func (x *TraceMessage) GetChain() *Chain {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *TraceMessage) GetTraces() []*Trace {
	if x != nil {
		return x.Traces
	}
	return nil
}

var File_ton_trace_message_proto protoreflect.FileDescriptor

var file_ton_trace_message_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x17, 0x74, 0x6f, 0x6e, 0x2f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc8, 0x01, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x70, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x08, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x52, 0x08, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52,
	0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x06, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ton_trace_message_proto_rawDescOnce sync.Once
	file_ton_trace_message_proto_rawDescData = file_ton_trace_message_proto_rawDesc
)

func file_ton_trace_message_proto_rawDescGZIP() []byte {
	file_ton_trace_message_proto_rawDescOnce.Do(func() {
		file_ton_trace_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_ton_trace_message_proto_rawDescData)
	})
	return file_ton_trace_message_proto_rawDescData
}

var file_ton_trace_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ton_trace_message_proto_goTypes = []any{
	(*Trace)(nil),        // 0: ton_messages.Trace
	(*TraceMessage)(nil), // 1: ton_messages.TraceMessage
	(*BlockHeader)(nil),  // 2: ton_messages.BlockHeader
	(*Transaction)(nil),  // 3: ton_messages.Transaction
	(*Chain)(nil),        // 4: ton_messages.Chain
}
var file_ton_trace_message_proto_depIdxs = []int32{
	2, // 0: ton_messages.Trace.BlockHeader:type_name -> ton_messages.BlockHeader
	3, // 1: ton_messages.Trace.Transaction:type_name -> ton_messages.Transaction
	0, // 2: ton_messages.Trace.Children:type_name -> ton_messages.Trace
	4, // 3: ton_messages.TraceMessage.Chain:type_name -> ton_messages.Chain
	0, // 4: ton_messages.TraceMessage.Traces:type_name -> ton_messages.Trace
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ton_trace_message_proto_init() }
func file_ton_trace_message_proto_init() {
	if File_ton_trace_message_proto != nil {
		return
	}
	file_ton_block_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ton_trace_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ton_trace_message_proto_goTypes,
		DependencyIndexes: file_ton_trace_message_proto_depIdxs,
		MessageInfos:      file_ton_trace_message_proto_msgTypes,
	}.Build()
	File_ton_trace_message_proto = out.File
	file_ton_trace_message_proto_rawDesc = nil
	file_ton_trace_message_proto_goTypes = nil
	file_ton_trace_message_proto_depIdxs = nil
}
