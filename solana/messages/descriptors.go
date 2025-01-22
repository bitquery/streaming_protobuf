package solana_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
)

type IndexRange [2]uint64
type Transactions struct {
	Timestamp, TransactionsCount, ProcessedCount uint64
	ExpectedCount                                *uint64
	IndexRanges                                  []IndexRange
}

type ExtendedBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	Transactions                               []Transactions
}

type ParsedIdlBlockMessageDescriptor struct {
	ExtendedBlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount              int
	InstructionsCount              int
	ParsedIdlInstructionsCount     int
}

type TokenBlockMessageDescriptor struct {
	ExtendedBlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount              int
	TransfersCount                 int
	TransactionBalanceUpdatesCount int
	InstructionBalanceUpdatesCount int
}

type DexBlockMessageDescriptor struct {
	ExtendedBlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount              int
	TradesCount                    int
	OrderEventsCount               int
	PoolChangeEventsCount          int
}
