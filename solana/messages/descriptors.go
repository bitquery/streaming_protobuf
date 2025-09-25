package solana_messages

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"

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

	if descriptor.Transactions == nil || len(descriptor.Transactions.IndexRanges) == 0 {
		return []byte(descriptor.BlockNumber.String())
	}

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

	// we need suffix as Kafka MurMur cache is sensitive to string suffixes, it will not distribute even partitions
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key[0:8], descriptor.BlockNumber.Uint64())
	binary.BigEndian.PutUint32(key[8:12], uint32(txIndex))

	h := fnv.New32a()
	h.Write(key)

	return []byte(fmt.Sprintf("%d-%d-%04X", descriptor.BlockNumber, txIndex, h.Sum32()))
}
