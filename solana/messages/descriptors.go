package solana_messages

import (
	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
)

type IndexRange [2]uint
type Transactions struct {
	Timestamp, TransactionsCount, ProcessedCount uint64
	ExpectedCount                                *uint64
	IndexRanges                                  []IndexRange
}

type ExtendedBlockMessageDescriptor struct {
	blockchain_messages.BlockMessageDescriptor `mapstructure:",squash"`
	Transactions                               *Transactions
}

const BitMaskForTxHashCorrelationId = 0b11
const BucketSizeForCorrelationId = 64

func (descriptor *ExtendedBlockMessageDescriptor) CorrelationId() []byte {
	return []byte(descriptor.BlockNumber.String())
	/*
		var txIndex = -1

		for _, indexRange := range descriptor.Transactions.IndexRanges {
			start := (indexRange[0] / BucketSizeForCorrelationId) & BitMaskForTxHashCorrelationId
			finish := (indexRange[1] / BucketSizeForCorrelationId) & BitMaskForTxHashCorrelationId

			if start != finish {
				return nil
			}

			if txIndex == -1 {
				txIndex = int(start)
				continue
			}

			if txIndex != int(start) {
				return nil
			}
		}

		if txIndex == -1 {
			return nil
		}
		return []byte(strconv.Itoa(txIndex))
	*/
}
