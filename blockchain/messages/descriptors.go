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

const BitMaskForTxHashCorrelationId = 0b11 // masks 2-bit field, 0 to 3 inclusive
var BitMaskCorrelationKeys = [][]byte{[]byte("red_stream"), []byte("blue_stream"), []byte("green_stream"), []byte("yellow_stream")}

func (descriptor *BlockMessageDescriptor) CorrelationId() []byte {
	if descriptor.IsBroadcasted() {

		if len(descriptor.TransactionHashes) == 1 {
			txHash := descriptor.TransactionHashes[0]
			return BitMaskCorrelationKeys[hexDigitToInt(txHash[len(txHash)-1])]
		}

		var txIndex = -1
		for _, txHash := range descriptor.TransactionHashes {
			index := hexDigitToInt(txHash[len(txHash)-1])
			if txIndex == -1 {
				txIndex = index
				continue
			}
			if txIndex != index {
				return nil
			}
		}
		if txIndex == -1 {
			return nil
		}
		return BitMaskCorrelationKeys[txIndex]
	}

	if len(descriptor.BlockHash) > 0 {
		return []byte(descriptor.BlockHash)
	}

	return nil

}

func hexDigitToInt(c byte) int {
	switch c {
	case '0', '4', '8', 'c', 'C':
		return 0
	case '1', '5', '9', 'd', 'D':
		return 1
	case '2', '6', 'a', 'A', 'e', 'E':
		return 2
	default:
		// Invalid hex digit — handle as needed
		return 3
	}
}
