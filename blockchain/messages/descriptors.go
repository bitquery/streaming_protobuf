package blockchain_messages

import (
	"github.com/bitquery/streaming_protobuf/pkg/util"
)

// BlockDescriptor is a base block definition structure
type BlockDescriptor struct {
	BlockHash    string
	BlockNumber  util.JsonNumber
	ParentHash   string
	ParentNumber util.JsonNumber
	Height       util.JsonNumber
}

// BlockMessageDescriptor is a descriptor for a block message, it may include transaction hashes for broadcasted
type BlockMessageDescriptor struct {
	BlockDescriptor   `mapstructure:",squash"`
	ChainId           string
	TransactionHashes []string `mapstructure:",omitempty" json:",omitempty"`
}
