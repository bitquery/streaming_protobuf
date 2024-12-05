// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: solana/token_block_message.proto

package solana_messages

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

type TokenCreator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Verified bool   `protobuf:"varint,2,opt,name=Verified,proto3" json:"Verified,omitempty"`
	Share    uint32 `protobuf:"varint,3,opt,name=Share,proto3" json:"Share,omitempty"`
}

func (x *TokenCreator) Reset() {
	*x = TokenCreator{}
	mi := &file_solana_token_block_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenCreator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCreator) ProtoMessage() {}

func (x *TokenCreator) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCreator.ProtoReflect.Descriptor instead.
func (*TokenCreator) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{0}
}

func (x *TokenCreator) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TokenCreator) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *TokenCreator) GetShare() uint32 {
	if x != nil {
		return x.Share
	}
	return 0
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string          `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Decimals             uint32          `protobuf:"varint,2,opt,name=Decimals,proto3" json:"Decimals,omitempty"`
	Uri                  string          `protobuf:"bytes,3,opt,name=Uri,proto3" json:"Uri,omitempty"`
	Symbol               string          `protobuf:"bytes,4,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Native               bool            `protobuf:"varint,5,opt,name=Native,proto3" json:"Native,omitempty"`
	Wrapped              bool            `protobuf:"varint,6,opt,name=Wrapped,proto3" json:"Wrapped,omitempty"`
	Fungible             bool            `protobuf:"varint,7,opt,name=Fungible,proto3" json:"Fungible,omitempty"`
	Key                  string          `protobuf:"bytes,8,opt,name=Key,proto3" json:"Key,omitempty"`
	SellerFeeBasisPoints uint32          `protobuf:"varint,9,opt,name=SellerFeeBasisPoints,proto3" json:"SellerFeeBasisPoints,omitempty"`
	EditionNonce         *uint32         `protobuf:"varint,10,opt,name=EditionNonce,proto3,oneof" json:"EditionNonce,omitempty"`
	TokenStandard        *string         `protobuf:"bytes,11,opt,name=TokenStandard,proto3,oneof" json:"TokenStandard,omitempty"`
	ProgramAddress       []byte          `protobuf:"bytes,12,opt,name=ProgramAddress,proto3" json:"ProgramAddress,omitempty"`
	MintAddress          []byte          `protobuf:"bytes,13,opt,name=MintAddress,proto3" json:"MintAddress,omitempty"`
	MetadataAddress      []byte          `protobuf:"bytes,14,opt,name=MetadataAddress,proto3" json:"MetadataAddress,omitempty"`
	UpdateAuthority      []byte          `protobuf:"bytes,15,opt,name=UpdateAuthority,proto3" json:"UpdateAuthority,omitempty"`
	CollectionAddress    []byte          `protobuf:"bytes,16,opt,name=CollectionAddress,proto3" json:"CollectionAddress,omitempty"`
	VerifiedCollection   bool            `protobuf:"varint,17,opt,name=VerifiedCollection,proto3" json:"VerifiedCollection,omitempty"`
	PrimarySaleHappened  bool            `protobuf:"varint,18,opt,name=PrimarySaleHappened,proto3" json:"PrimarySaleHappened,omitempty"`
	IsMutable            bool            `protobuf:"varint,19,opt,name=IsMutable,proto3" json:"IsMutable,omitempty"`
	TokenCreators        []*TokenCreator `protobuf:"bytes,20,rep,name=TokenCreators,proto3" json:"TokenCreators,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	mi := &file_solana_token_block_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Currency) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *Currency) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Currency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Currency) GetNative() bool {
	if x != nil {
		return x.Native
	}
	return false
}

func (x *Currency) GetWrapped() bool {
	if x != nil {
		return x.Wrapped
	}
	return false
}

func (x *Currency) GetFungible() bool {
	if x != nil {
		return x.Fungible
	}
	return false
}

func (x *Currency) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Currency) GetSellerFeeBasisPoints() uint32 {
	if x != nil {
		return x.SellerFeeBasisPoints
	}
	return 0
}

func (x *Currency) GetEditionNonce() uint32 {
	if x != nil && x.EditionNonce != nil {
		return *x.EditionNonce
	}
	return 0
}

func (x *Currency) GetTokenStandard() string {
	if x != nil && x.TokenStandard != nil {
		return *x.TokenStandard
	}
	return ""
}

func (x *Currency) GetProgramAddress() []byte {
	if x != nil {
		return x.ProgramAddress
	}
	return nil
}

func (x *Currency) GetMintAddress() []byte {
	if x != nil {
		return x.MintAddress
	}
	return nil
}

func (x *Currency) GetMetadataAddress() []byte {
	if x != nil {
		return x.MetadataAddress
	}
	return nil
}

func (x *Currency) GetUpdateAuthority() []byte {
	if x != nil {
		return x.UpdateAuthority
	}
	return nil
}

func (x *Currency) GetCollectionAddress() []byte {
	if x != nil {
		return x.CollectionAddress
	}
	return nil
}

func (x *Currency) GetVerifiedCollection() bool {
	if x != nil {
		return x.VerifiedCollection
	}
	return false
}

func (x *Currency) GetPrimarySaleHappened() bool {
	if x != nil {
		return x.PrimarySaleHappened
	}
	return false
}

func (x *Currency) GetIsMutable() bool {
	if x != nil {
		return x.IsMutable
	}
	return false
}

func (x *Currency) GetTokenCreators() []*TokenCreator {
	if x != nil {
		return x.TokenCreators
	}
	return nil
}

type CurrencyBalanceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BalanceUpdate *BalanceUpdate `protobuf:"bytes,1,opt,name=BalanceUpdate,proto3" json:"BalanceUpdate,omitempty"`
	Currency      *Currency      `protobuf:"bytes,2,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *CurrencyBalanceUpdate) Reset() {
	*x = CurrencyBalanceUpdate{}
	mi := &file_solana_token_block_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyBalanceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyBalanceUpdate) ProtoMessage() {}

func (x *CurrencyBalanceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyBalanceUpdate.ProtoReflect.Descriptor instead.
func (*CurrencyBalanceUpdate) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyBalanceUpdate) GetBalanceUpdate() *BalanceUpdate {
	if x != nil {
		return x.BalanceUpdate
	}
	return nil
}

func (x *CurrencyBalanceUpdate) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstructionIndex uint32    `protobuf:"varint,1,opt,name=InstructionIndex,proto3" json:"InstructionIndex,omitempty"`
	Amount           uint64    `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Sender           *Account  `protobuf:"bytes,3,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Receiver         *Account  `protobuf:"bytes,4,opt,name=Receiver,proto3" json:"Receiver,omitempty"`
	Authority        *Account  `protobuf:"bytes,5,opt,name=Authority,proto3" json:"Authority,omitempty"`
	Currency         *Currency `protobuf:"bytes,6,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	mi := &file_solana_token_block_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{3}
}

func (x *Transfer) GetInstructionIndex() uint32 {
	if x != nil {
		return x.InstructionIndex
	}
	return 0
}

func (x *Transfer) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transfer) GetSender() *Account {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Transfer) GetReceiver() *Account {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Transfer) GetAuthority() *Account {
	if x != nil {
		return x.Authority
	}
	return nil
}

func (x *Transfer) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

type InstructionBalanceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstructionIndex            uint32                   `protobuf:"varint,1,opt,name=InstructionIndex,proto3" json:"InstructionIndex,omitempty"`
	TotalCurrencyBalanceUpdates []*CurrencyBalanceUpdate `protobuf:"bytes,2,rep,name=TotalCurrencyBalanceUpdates,proto3" json:"TotalCurrencyBalanceUpdates,omitempty"`
	OwnCurrencyBalanceUpdates   []*CurrencyBalanceUpdate `protobuf:"bytes,3,rep,name=OwnCurrencyBalanceUpdates,proto3" json:"OwnCurrencyBalanceUpdates,omitempty"`
}

func (x *InstructionBalanceUpdate) Reset() {
	*x = InstructionBalanceUpdate{}
	mi := &file_solana_token_block_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstructionBalanceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstructionBalanceUpdate) ProtoMessage() {}

func (x *InstructionBalanceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstructionBalanceUpdate.ProtoReflect.Descriptor instead.
func (*InstructionBalanceUpdate) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{4}
}

func (x *InstructionBalanceUpdate) GetInstructionIndex() uint32 {
	if x != nil {
		return x.InstructionIndex
	}
	return 0
}

func (x *InstructionBalanceUpdate) GetTotalCurrencyBalanceUpdates() []*CurrencyBalanceUpdate {
	if x != nil {
		return x.TotalCurrencyBalanceUpdates
	}
	return nil
}

func (x *InstructionBalanceUpdate) GetOwnCurrencyBalanceUpdates() []*CurrencyBalanceUpdate {
	if x != nil {
		return x.OwnCurrencyBalanceUpdates
	}
	return nil
}

type ParsedTokenTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index                     uint32                      `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Signature                 []byte                      `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Status                    *TransactionStatus          `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Header                    *TransactionHeader          `protobuf:"bytes,4,opt,name=Header,proto3" json:"Header,omitempty"`
	Transfers                 []*Transfer                 `protobuf:"bytes,5,rep,name=Transfers,proto3" json:"Transfers,omitempty"`
	BalanceUpdates            []*CurrencyBalanceUpdate    `protobuf:"bytes,6,rep,name=BalanceUpdates,proto3" json:"BalanceUpdates,omitempty"`
	InstructionBalanceUpdates []*InstructionBalanceUpdate `protobuf:"bytes,7,rep,name=InstructionBalanceUpdates,proto3" json:"InstructionBalanceUpdates,omitempty"`
}

func (x *ParsedTokenTransaction) Reset() {
	*x = ParsedTokenTransaction{}
	mi := &file_solana_token_block_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParsedTokenTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedTokenTransaction) ProtoMessage() {}

func (x *ParsedTokenTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedTokenTransaction.ProtoReflect.Descriptor instead.
func (*ParsedTokenTransaction) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{5}
}

func (x *ParsedTokenTransaction) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ParsedTokenTransaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *ParsedTokenTransaction) GetStatus() *TransactionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ParsedTokenTransaction) GetHeader() *TransactionHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ParsedTokenTransaction) GetTransfers() []*Transfer {
	if x != nil {
		return x.Transfers
	}
	return nil
}

func (x *ParsedTokenTransaction) GetBalanceUpdates() []*CurrencyBalanceUpdate {
	if x != nil {
		return x.BalanceUpdates
	}
	return nil
}

func (x *ParsedTokenTransaction) GetInstructionBalanceUpdates() []*InstructionBalanceUpdate {
	if x != nil {
		return x.InstructionBalanceUpdates
	}
	return nil
}

type TokenBlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader              `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Transactions []*ParsedTokenTransaction `protobuf:"bytes,2,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
}

func (x *TokenBlockMessage) Reset() {
	*x = TokenBlockMessage{}
	mi := &file_solana_token_block_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenBlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBlockMessage) ProtoMessage() {}

func (x *TokenBlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_solana_token_block_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBlockMessage.ProtoReflect.Descriptor instead.
func (*TokenBlockMessage) Descriptor() ([]byte, []int) {
	return file_solana_token_block_message_proto_rawDescGZIP(), []int{6}
}

func (x *TokenBlockMessage) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenBlockMessage) GetTransactions() []*ParsedTokenTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_solana_token_block_message_proto protoreflect.FileDescriptor

var file_solana_token_block_message_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x1a, 0x1a, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x22, 0x80, 0x06, 0x0a, 0x08,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x69, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x67, 0x69, 0x62, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x46, 0x75, 0x6e, 0x67, 0x69, 0x62, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x65, 0x65, 0x42,
	0x61, 0x73, 0x69, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x14, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x65, 0x65, 0x42, 0x61, 0x73, 0x69, 0x73,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0c,
	0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x29, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x4d, 0x69, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x61, 0x6c, 0x65, 0x48, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x61, 0x6c, 0x65,
	0x48, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x4d, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x4d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x22, 0x94,
	0x01, 0x0a, 0x15, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x0d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xa5, 0x02, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6f, 0x6c,
	0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x36,
	0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e,
	0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x96, 0x02,
	0x0a, 0x18, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x68, 0x0a, 0x1b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f,
	0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x1b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x64, 0x0a, 0x19, 0x4f, 0x77, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x19, 0x4f, 0x77, 0x6e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0xb6, 0x03, 0x0a, 0x16, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x37, 0x0a,
	0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x09, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6f, 0x6c, 0x61,
	0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x96, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_solana_token_block_message_proto_rawDescOnce sync.Once
	file_solana_token_block_message_proto_rawDescData = file_solana_token_block_message_proto_rawDesc
)

func file_solana_token_block_message_proto_rawDescGZIP() []byte {
	file_solana_token_block_message_proto_rawDescOnce.Do(func() {
		file_solana_token_block_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_solana_token_block_message_proto_rawDescData)
	})
	return file_solana_token_block_message_proto_rawDescData
}

var file_solana_token_block_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_solana_token_block_message_proto_goTypes = []any{
	(*TokenCreator)(nil),             // 0: solana_messages.TokenCreator
	(*Currency)(nil),                 // 1: solana_messages.Currency
	(*CurrencyBalanceUpdate)(nil),    // 2: solana_messages.CurrencyBalanceUpdate
	(*Transfer)(nil),                 // 3: solana_messages.Transfer
	(*InstructionBalanceUpdate)(nil), // 4: solana_messages.InstructionBalanceUpdate
	(*ParsedTokenTransaction)(nil),   // 5: solana_messages.ParsedTokenTransaction
	(*TokenBlockMessage)(nil),        // 6: solana_messages.TokenBlockMessage
	(*BalanceUpdate)(nil),            // 7: solana_messages.BalanceUpdate
	(*Account)(nil),                  // 8: solana_messages.Account
	(*TransactionStatus)(nil),        // 9: solana_messages.TransactionStatus
	(*TransactionHeader)(nil),        // 10: solana_messages.TransactionHeader
	(*BlockHeader)(nil),              // 11: solana_messages.BlockHeader
}
var file_solana_token_block_message_proto_depIdxs = []int32{
	0,  // 0: solana_messages.Currency.TokenCreators:type_name -> solana_messages.TokenCreator
	7,  // 1: solana_messages.CurrencyBalanceUpdate.BalanceUpdate:type_name -> solana_messages.BalanceUpdate
	1,  // 2: solana_messages.CurrencyBalanceUpdate.Currency:type_name -> solana_messages.Currency
	8,  // 3: solana_messages.Transfer.Sender:type_name -> solana_messages.Account
	8,  // 4: solana_messages.Transfer.Receiver:type_name -> solana_messages.Account
	8,  // 5: solana_messages.Transfer.Authority:type_name -> solana_messages.Account
	1,  // 6: solana_messages.Transfer.Currency:type_name -> solana_messages.Currency
	2,  // 7: solana_messages.InstructionBalanceUpdate.TotalCurrencyBalanceUpdates:type_name -> solana_messages.CurrencyBalanceUpdate
	2,  // 8: solana_messages.InstructionBalanceUpdate.OwnCurrencyBalanceUpdates:type_name -> solana_messages.CurrencyBalanceUpdate
	9,  // 9: solana_messages.ParsedTokenTransaction.Status:type_name -> solana_messages.TransactionStatus
	10, // 10: solana_messages.ParsedTokenTransaction.Header:type_name -> solana_messages.TransactionHeader
	3,  // 11: solana_messages.ParsedTokenTransaction.Transfers:type_name -> solana_messages.Transfer
	2,  // 12: solana_messages.ParsedTokenTransaction.BalanceUpdates:type_name -> solana_messages.CurrencyBalanceUpdate
	4,  // 13: solana_messages.ParsedTokenTransaction.InstructionBalanceUpdates:type_name -> solana_messages.InstructionBalanceUpdate
	11, // 14: solana_messages.TokenBlockMessage.Header:type_name -> solana_messages.BlockHeader
	5,  // 15: solana_messages.TokenBlockMessage.Transactions:type_name -> solana_messages.ParsedTokenTransaction
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_solana_token_block_message_proto_init() }
func file_solana_token_block_message_proto_init() {
	if File_solana_token_block_message_proto != nil {
		return
	}
	file_solana_block_message_proto_init()
	file_solana_token_block_message_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_solana_token_block_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_solana_token_block_message_proto_goTypes,
		DependencyIndexes: file_solana_token_block_message_proto_depIdxs,
		MessageInfos:      file_solana_token_block_message_proto_msgTypes,
	}.Build()
	File_solana_token_block_message_proto = out.File
	file_solana_token_block_message_proto_rawDesc = nil
	file_solana_token_block_message_proto_goTypes = nil
	file_solana_token_block_message_proto_depIdxs = nil
}
