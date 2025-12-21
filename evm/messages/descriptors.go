package evm_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
)

type DexBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TradesCount                                uint32
}

type DexPoolMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	PoolEventsCount                            uint32
}

type ParsedAbiBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	MinerReward                                string
	UncleRewardsCount                          uint32
	TransactionsCount                          uint32
}

type TokenBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TransfersCount                             uint32
}
