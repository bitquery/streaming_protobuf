package solana_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/blockchain/messages"
)

type ParsedIdlBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount                          int
	InstructionsCount                          int
	ParsedIdlInstructionsCount                 int
}

type TokenBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount                          int
	TransfersCount                             int
	TransactionBalanceUpdatesCount             int
	InstructionBalanceUpdatesCount             int
}

type DexBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount                          int
	TradesCount                                int
	OrderEventsCount                           int
	PoolChangeEventsCount                      int
}
