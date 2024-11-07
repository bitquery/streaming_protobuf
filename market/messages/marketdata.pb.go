// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: market/marketdata.proto

package marketdata_messages

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

type MarketdataMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currencies []*Currency `protobuf:"bytes,1,rep,name=Currencies,proto3" json:"Currencies,omitempty"`
	Source     string      `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"` // cryptorank for example
}

func (x *MarketdataMessage) Reset() {
	*x = MarketdataMessage{}
	mi := &file_market_marketdata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MarketdataMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketdataMessage) ProtoMessage() {}

func (x *MarketdataMessage) ProtoReflect() protoreflect.Message {
	mi := &file_market_marketdata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketdataMessage.ProtoReflect.Descriptor instead.
func (*MarketdataMessage) Descriptor() ([]byte, []int) {
	return file_market_marketdata_proto_rawDescGZIP(), []int{0}
}

func (x *MarketdataMessage) GetCurrencies() []*Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

func (x *MarketdataMessage) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Did      string  `protobuf:"bytes,2,opt,name=Did,proto3" json:"Did,omitempty"`
	Network  string  `protobuf:"bytes,3,opt,name=Network,proto3" json:"Network,omitempty"`
	Address  *string `protobuf:"bytes,4,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	TokenId  *string `protobuf:"bytes,5,opt,name=TokenId,proto3,oneof" json:"TokenId,omitempty"` // For networks like TRON and TRC-10
	IsNative bool    `protobuf:"varint,6,opt,name=IsNative,proto3" json:"IsNative,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_market_marketdata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_market_marketdata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_market_marketdata_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Token) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *Token) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Token) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *Token) GetTokenId() string {
	if x != nil && x.TokenId != nil {
		return *x.TokenId
	}
	return ""
}

func (x *Token) GetIsNative() bool {
	if x != nil {
		return x.IsNative
	}
	return false
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens     []*Token     `protobuf:"bytes,1,rep,name=Tokens,proto3" json:"Tokens,omitempty"`
	Name       string       `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Symbol     string       `protobuf:"bytes,3,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	OriginalId string       `protobuf:"bytes,4,opt,name=OriginalId,proto3" json:"OriginalId,omitempty"`
	Bucket     *PriceBucket `protobuf:"bytes,6,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	mi := &file_market_marketdata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_market_marketdata_proto_msgTypes[2]
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
	return file_market_marketdata_proto_rawDescGZIP(), []int{2}
}

func (x *Currency) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Currency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Currency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Currency) GetOriginalId() string {
	if x != nil {
		return x.OriginalId
	}
	return ""
}

func (x *Currency) GetBucket() *PriceBucket {
	if x != nil {
		return x.Bucket
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time uint64  `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	Usd  float64 `protobuf:"fixed64,2,opt,name=Usd,proto3" json:"Usd,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	mi := &file_market_marketdata_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_market_marketdata_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_market_marketdata_proto_rawDescGZIP(), []int{3}
}

func (x *Price) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Price) GetUsd() float64 {
	if x != nil {
		return x.Usd
	}
	return 0
}

type PriceBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval      uint32   `protobuf:"varint,1,opt,name=Interval,proto3" json:"Interval,omitempty"` // Interval in seconds between prices
	Prices        []*Price `protobuf:"bytes,2,rep,name=Prices,proto3" json:"Prices,omitempty"`
	LastUpdatedAt uint64   `protobuf:"varint,3,opt,name=LastUpdatedAt,proto3" json:"LastUpdatedAt,omitempty"`
}

func (x *PriceBucket) Reset() {
	*x = PriceBucket{}
	mi := &file_market_marketdata_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PriceBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceBucket) ProtoMessage() {}

func (x *PriceBucket) ProtoReflect() protoreflect.Message {
	mi := &file_market_marketdata_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceBucket.ProtoReflect.Descriptor instead.
func (*PriceBucket) Descriptor() ([]byte, []int) {
	return file_market_marketdata_proto_rawDescGZIP(), []int{4}
}

func (x *PriceBucket) GetInterval() uint32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *PriceBucket) GetPrices() []*Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *PriceBucket) GetLastUpdatedAt() uint64 {
	if x != nil {
		return x.LastUpdatedAt
	}
	return 0
}

var File_market_marketdata_proto protoreflect.FileDescriptor

var file_market_marketdata_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x6a,
	0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x44, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x1d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x07, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x73, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x49, 0x73, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x32, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x73, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x55, 0x73, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_market_marketdata_proto_rawDescOnce sync.Once
	file_market_marketdata_proto_rawDescData = file_market_marketdata_proto_rawDesc
)

func file_market_marketdata_proto_rawDescGZIP() []byte {
	file_market_marketdata_proto_rawDescOnce.Do(func() {
		file_market_marketdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_marketdata_proto_rawDescData)
	})
	return file_market_marketdata_proto_rawDescData
}

var file_market_marketdata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_market_marketdata_proto_goTypes = []any{
	(*MarketdataMessage)(nil), // 0: marketdata_messages.MarketdataMessage
	(*Token)(nil),             // 1: marketdata_messages.Token
	(*Currency)(nil),          // 2: marketdata_messages.Currency
	(*Price)(nil),             // 3: marketdata_messages.Price
	(*PriceBucket)(nil),       // 4: marketdata_messages.PriceBucket
}
var file_market_marketdata_proto_depIdxs = []int32{
	2, // 0: marketdata_messages.MarketdataMessage.Currencies:type_name -> marketdata_messages.Currency
	1, // 1: marketdata_messages.Currency.Tokens:type_name -> marketdata_messages.Token
	4, // 2: marketdata_messages.Currency.Bucket:type_name -> marketdata_messages.PriceBucket
	3, // 3: marketdata_messages.PriceBucket.Prices:type_name -> marketdata_messages.Price
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_market_marketdata_proto_init() }
func file_market_marketdata_proto_init() {
	if File_market_marketdata_proto != nil {
		return
	}
	file_market_marketdata_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_marketdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_market_marketdata_proto_goTypes,
		DependencyIndexes: file_market_marketdata_proto_depIdxs,
		MessageInfos:      file_market_marketdata_proto_msgTypes,
	}.Build()
	File_market_marketdata_proto = out.File
	file_market_marketdata_proto_rawDesc = nil
	file_market_marketdata_proto_goTypes = nil
	file_market_marketdata_proto_depIdxs = nil
}
