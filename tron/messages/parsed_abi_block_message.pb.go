// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: tron/parsed_abi_block_message.proto

package tron_messages

import (
	messages "evm/messages"
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

type ContractInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Type    string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	TypeUrl string `protobuf:"bytes,3,opt,name=TypeUrl,proto3" json:"TypeUrl,omitempty"`
}

func (x *ContractInfo) Reset() {
	*x = ContractInfo{}
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContractInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractInfo) ProtoMessage() {}

func (x *ContractInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractInfo.ProtoReflect.Descriptor instead.
func (*ContractInfo) Descriptor() ([]byte, []int) {
	return file_tron_parsed_abi_block_message_proto_rawDescGZIP(), []int{0}
}

func (x *ContractInfo) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ContractInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ContractInfo) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

type ParsedAbiTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHeader *TransactionHeader                   `protobuf:"bytes,1,opt,name=TransactionHeader,proto3" json:"TransactionHeader,omitempty"`
	Receipt           *Receipt                             `protobuf:"bytes,2,opt,name=Receipt,proto3" json:"Receipt,omitempty"`
	TransactionStatus *messages.ParsedAbiTransactionStatus `protobuf:"bytes,3,opt,name=TransactionStatus,proto3" json:"TransactionStatus,omitempty"`
	TransactionResult *TransactionResult                   `protobuf:"bytes,4,opt,name=TransactionResult,proto3" json:"TransactionResult,omitempty"`
	ContractInfo      *ContractInfo                        `protobuf:"bytes,5,opt,name=ContractInfo,proto3" json:"ContractInfo,omitempty"`
	Calls             []*messages.ParsedAbiCall            `protobuf:"bytes,6,rep,name=Calls,proto3" json:"Calls,omitempty"`
	RewardWithdraw    *RewardWithdraw                      `protobuf:"bytes,7,opt,name=RewardWithdraw,proto3" json:"RewardWithdraw,omitempty"`
}

func (x *ParsedAbiTransaction) Reset() {
	*x = ParsedAbiTransaction{}
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParsedAbiTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedAbiTransaction) ProtoMessage() {}

func (x *ParsedAbiTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedAbiTransaction.ProtoReflect.Descriptor instead.
func (*ParsedAbiTransaction) Descriptor() ([]byte, []int) {
	return file_tron_parsed_abi_block_message_proto_rawDescGZIP(), []int{1}
}

func (x *ParsedAbiTransaction) GetTransactionHeader() *TransactionHeader {
	if x != nil {
		return x.TransactionHeader
	}
	return nil
}

func (x *ParsedAbiTransaction) GetReceipt() *Receipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

func (x *ParsedAbiTransaction) GetTransactionStatus() *messages.ParsedAbiTransactionStatus {
	if x != nil {
		return x.TransactionStatus
	}
	return nil
}

func (x *ParsedAbiTransaction) GetTransactionResult() *TransactionResult {
	if x != nil {
		return x.TransactionResult
	}
	return nil
}

func (x *ParsedAbiTransaction) GetContractInfo() *ContractInfo {
	if x != nil {
		return x.ContractInfo
	}
	return nil
}

func (x *ParsedAbiTransaction) GetCalls() []*messages.ParsedAbiCall {
	if x != nil {
		return x.Calls
	}
	return nil
}

func (x *ParsedAbiTransaction) GetRewardWithdraw() *RewardWithdraw {
	if x != nil {
		return x.RewardWithdraw
	}
	return nil
}

type ParsedAbiBlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain        *Chain                  `protobuf:"bytes,1,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Header       *BlockHeader            `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	Witness      *Witness                `protobuf:"bytes,3,opt,name=Witness,proto3" json:"Witness,omitempty"`
	Transactions []*ParsedAbiTransaction `protobuf:"bytes,4,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
}

func (x *ParsedAbiBlockMessage) Reset() {
	*x = ParsedAbiBlockMessage{}
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParsedAbiBlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedAbiBlockMessage) ProtoMessage() {}

func (x *ParsedAbiBlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tron_parsed_abi_block_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedAbiBlockMessage.ProtoReflect.Descriptor instead.
func (*ParsedAbiBlockMessage) Descriptor() ([]byte, []int) {
	return file_tron_parsed_abi_block_message_proto_rawDescGZIP(), []int{2}
}

func (x *ParsedAbiBlockMessage) GetChain() *Chain {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *ParsedAbiBlockMessage) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ParsedAbiBlockMessage) GetWitness() *Witness {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *ParsedAbiBlockMessage) GetTransactions() []*ParsedAbiTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_tron_parsed_abi_block_message_proto protoreflect.FileDescriptor

var file_tron_parsed_abi_block_message_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x62,
	0x69, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x1a, 0x18, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x65, 0x76, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x62, 0x69, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xfb, 0x03, 0x0a, 0x14, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x64, 0x41, 0x62, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x56, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x65, 0x76, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x41, 0x62, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31,
	0x0a, 0x05, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x65, 0x76, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x64, 0x41, 0x62, 0x69, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x43, 0x61, 0x6c, 0x6c,
	0x73, 0x12, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x72, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x22, 0xf2, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x64, 0x41, 0x62, 0x69, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x32,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x07, 0x57, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x64, 0x41, 0x62, 0x69, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tron_parsed_abi_block_message_proto_rawDescOnce sync.Once
	file_tron_parsed_abi_block_message_proto_rawDescData = file_tron_parsed_abi_block_message_proto_rawDesc
)

func file_tron_parsed_abi_block_message_proto_rawDescGZIP() []byte {
	file_tron_parsed_abi_block_message_proto_rawDescOnce.Do(func() {
		file_tron_parsed_abi_block_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_tron_parsed_abi_block_message_proto_rawDescData)
	})
	return file_tron_parsed_abi_block_message_proto_rawDescData
}

var file_tron_parsed_abi_block_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tron_parsed_abi_block_message_proto_goTypes = []any{
	(*ContractInfo)(nil),                        // 0: tron_messages.ContractInfo
	(*ParsedAbiTransaction)(nil),                // 1: tron_messages.ParsedAbiTransaction
	(*ParsedAbiBlockMessage)(nil),               // 2: tron_messages.ParsedAbiBlockMessage
	(*TransactionHeader)(nil),                   // 3: tron_messages.TransactionHeader
	(*Receipt)(nil),                             // 4: tron_messages.Receipt
	(*messages.ParsedAbiTransactionStatus)(nil), // 5: evm_messages.ParsedAbiTransactionStatus
	(*TransactionResult)(nil),                   // 6: tron_messages.TransactionResult
	(*messages.ParsedAbiCall)(nil),              // 7: evm_messages.ParsedAbiCall
	(*RewardWithdraw)(nil),                      // 8: tron_messages.RewardWithdraw
	(*Chain)(nil),                               // 9: tron_messages.Chain
	(*BlockHeader)(nil),                         // 10: tron_messages.BlockHeader
	(*Witness)(nil),                             // 11: tron_messages.Witness
}
var file_tron_parsed_abi_block_message_proto_depIdxs = []int32{
	3,  // 0: tron_messages.ParsedAbiTransaction.TransactionHeader:type_name -> tron_messages.TransactionHeader
	4,  // 1: tron_messages.ParsedAbiTransaction.Receipt:type_name -> tron_messages.Receipt
	5,  // 2: tron_messages.ParsedAbiTransaction.TransactionStatus:type_name -> evm_messages.ParsedAbiTransactionStatus
	6,  // 3: tron_messages.ParsedAbiTransaction.TransactionResult:type_name -> tron_messages.TransactionResult
	0,  // 4: tron_messages.ParsedAbiTransaction.ContractInfo:type_name -> tron_messages.ContractInfo
	7,  // 5: tron_messages.ParsedAbiTransaction.Calls:type_name -> evm_messages.ParsedAbiCall
	8,  // 6: tron_messages.ParsedAbiTransaction.RewardWithdraw:type_name -> tron_messages.RewardWithdraw
	9,  // 7: tron_messages.ParsedAbiBlockMessage.Chain:type_name -> tron_messages.Chain
	10, // 8: tron_messages.ParsedAbiBlockMessage.Header:type_name -> tron_messages.BlockHeader
	11, // 9: tron_messages.ParsedAbiBlockMessage.Witness:type_name -> tron_messages.Witness
	1,  // 10: tron_messages.ParsedAbiBlockMessage.Transactions:type_name -> tron_messages.ParsedAbiTransaction
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tron_parsed_abi_block_message_proto_init() }
func file_tron_parsed_abi_block_message_proto_init() {
	if File_tron_parsed_abi_block_message_proto != nil {
		return
	}
	file_tron_block_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tron_parsed_abi_block_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tron_parsed_abi_block_message_proto_goTypes,
		DependencyIndexes: file_tron_parsed_abi_block_message_proto_depIdxs,
		MessageInfos:      file_tron_parsed_abi_block_message_proto_msgTypes,
	}.Build()
	File_tron_parsed_abi_block_message_proto = out.File
	file_tron_parsed_abi_block_message_proto_rawDesc = nil
	file_tron_parsed_abi_block_message_proto_goTypes = nil
	file_tron_parsed_abi_block_message_proto_depIdxs = nil
}
