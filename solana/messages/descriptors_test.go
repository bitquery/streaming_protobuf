package solana_messages

import (
	"testing"

	blockchain_messages "github.com/bitquery/streaming_protobuf/v2/blockchain/messages"
	"github.com/stretchr/testify/assert"
)

func TestOneCorrId(t *testing.T) {

	t.Parallel()

	descriptor := &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{0, 1},
			},
		},
	}

	zero := descriptor.CorrelationId()
	assert.Equal(t, "12-0", string(descriptor.CorrelationId())[0:4])

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
			},
		},
	}

	assert.Equal(t, "12-1", string(descriptor.CorrelationId())[0:4])

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1000, 1004},
			},
		},
	}

	assert.Equal(t, "12-3", string(descriptor.CorrelationId())[0:4])

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1028, 1029},
			},
		},
	}

	assert.Equal(t, "12-0", string(descriptor.CorrelationId())[0:4])
	assert.Equal(t, string(zero), string(descriptor.CorrelationId()))

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1028, 1129},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

}

func TestManyCorrId(t *testing.T) {
	t.Parallel()

	descriptor := &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{0, 1},
				{10, 11},
			},
		},
	}

	assert.Equal(t, "12-0", string(descriptor.CorrelationId())[0:4])

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1024 + 100, 1024 + 104},
			},
		},
	}

	assert.Equal(t, "12-1", string(descriptor.CorrelationId())[0:4])

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1028, 1129},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		BlockMessageDescriptor: blockchain_messages.BlockMessageDescriptor{
			BlockDescriptor: blockchain_messages.BlockDescriptor{
				BlockNumber: 12,
			},
		},
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1028, 1029},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

}
