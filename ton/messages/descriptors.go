package ton_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
)

type TonSingleBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	ShardId                                    string
	WorkchainId                                int32
}

type TonBlockMessageDescriptor struct {
	MinBlock TonSingleBlockMessageDescriptor
	MaxBlock TonSingleBlockMessageDescriptor
	Blocks   []TonSingleBlockMessageDescriptor
}

type TokenBlockMessageDescriptor struct {
	TonBlockMessageDescriptor `mapstructure:",squash"`
	TransfersCount            uint32
}

type DexBlockMessageDescriptor struct {
	TonBlockMessageDescriptor `mapstructure:",squash"`
	TradesCount               uint32
}
