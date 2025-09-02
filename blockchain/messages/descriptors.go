package blockchain_messages

import (
	"github.com/bitquery/streaming_protobuf/v2/pkg/encoder"
)

// BlockDescriptor is a base block definition structure
type BlockDescriptor struct {
	BlockHash    string
	BlockNumber  encoder.JsonNumber
	ParentHash   string
	ParentNumber encoder.JsonNumber
	Height       encoder.JsonNumber
}

// BlockMessageDescriptor is a descriptor for a block message, it may include transaction hashes for broadcasted
type BlockMessageDescriptor struct {
	BlockDescriptor   `mapstructure:",squash"`
	ChainId           string
	TransactionHashes []string `mapstructure:",omitempty" json:",omitempty"`
}

/*
For Solana height is not equal to block number
For EVM blockchains they are the same and we do not send Height in descriptor
*/
func (d BlockDescriptor) BlockHeight() int64 {
	if d.Height != 0 {
		return int64(d.Height)
	}
	return int64(d.BlockNumber)
}

func (descriptor *BlockMessageDescriptor) IsBroadcasted() bool {
	return len(descriptor.TransactionHashes) > 0
}

func (descriptor *BlockMessageDescriptor) CorrelationId() []byte {
	if descriptor.IsBroadcasted() || len(descriptor.BlockHash) == 0 {
		return nil
	}
	return []byte(descriptor.BlockHash)
}
