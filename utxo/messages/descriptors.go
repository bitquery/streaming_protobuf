package utxo_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
)

type ParsedBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	TransactionsCount                          uint32
	InputsCount                                uint32
	OutputsCount                               uint32
}
