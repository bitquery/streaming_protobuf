// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: tron/dex_block_message.proto

package tron_messages

import (
	messages "github.com/bitquery/streaming_protobuf/evm/messages"
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

type DexBlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain  *Chain               `protobuf:"bytes,1,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Header *BlockHeader         `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	Trades []*messages.DexTrade `protobuf:"bytes,3,rep,name=Trades,proto3" json:"Trades,omitempty"`
}

func (x *DexBlockMessage) Reset() {
	*x = DexBlockMessage{}
	mi := &file_tron_dex_block_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DexBlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DexBlockMessage) ProtoMessage() {}

func (x *DexBlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tron_dex_block_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DexBlockMessage.ProtoReflect.Descriptor instead.
func (*DexBlockMessage) Descriptor() ([]byte, []int) {
	return file_tron_dex_block_message_proto_rawDescGZIP(), []int{0}
}

func (x *DexBlockMessage) GetChain() *Chain {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *DexBlockMessage) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DexBlockMessage) GetTrades() []*messages.DexTrade {
	if x != nil {
		return x.Trades
	}
	return nil
}

var File_tron_dex_block_message_proto protoreflect.FileDescriptor

var file_tron_dex_block_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x18, 0x74,
	0x72, 0x6f, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x76, 0x6d, 0x2f, 0x64, 0x65, 0x78,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x78, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x6d, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x78, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tron_dex_block_message_proto_rawDescOnce sync.Once
	file_tron_dex_block_message_proto_rawDescData = file_tron_dex_block_message_proto_rawDesc
)

func file_tron_dex_block_message_proto_rawDescGZIP() []byte {
	file_tron_dex_block_message_proto_rawDescOnce.Do(func() {
		file_tron_dex_block_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_tron_dex_block_message_proto_rawDescData)
	})
	return file_tron_dex_block_message_proto_rawDescData
}

var file_tron_dex_block_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tron_dex_block_message_proto_goTypes = []any{
	(*DexBlockMessage)(nil),   // 0: tron_messages.DexBlockMessage
	(*Chain)(nil),             // 1: tron_messages.Chain
	(*BlockHeader)(nil),       // 2: tron_messages.BlockHeader
	(*messages.DexTrade)(nil), // 3: evm_messages.DexTrade
}
var file_tron_dex_block_message_proto_depIdxs = []int32{
	1, // 0: tron_messages.DexBlockMessage.Chain:type_name -> tron_messages.Chain
	2, // 1: tron_messages.DexBlockMessage.Header:type_name -> tron_messages.BlockHeader
	3, // 2: tron_messages.DexBlockMessage.Trades:type_name -> evm_messages.DexTrade
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tron_dex_block_message_proto_init() }
func file_tron_dex_block_message_proto_init() {
	if File_tron_dex_block_message_proto != nil {
		return
	}
	file_tron_block_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tron_dex_block_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tron_dex_block_message_proto_goTypes,
		DependencyIndexes: file_tron_dex_block_message_proto_depIdxs,
		MessageInfos:      file_tron_dex_block_message_proto_msgTypes,
	}.Build()
	File_tron_dex_block_message_proto = out.File
	file_tron_dex_block_message_proto_rawDesc = nil
	file_tron_dex_block_message_proto_goTypes = nil
	file_tron_dex_block_message_proto_depIdxs = nil
}