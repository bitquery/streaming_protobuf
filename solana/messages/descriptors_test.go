package solana_messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneCorrId(t *testing.T) {
	t.Skip()
	t.Parallel()

	descriptor := &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{0, 1},
			},
		},
	}

	assert.Equal(t, []byte("0"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
			},
		},
	}

	assert.Equal(t, []byte("1"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1000, 1004},
			},
		},
	}

	assert.Equal(t, []byte("3"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1028, 1029},
			},
		},
	}

	assert.Equal(t, []byte("0"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{1028, 1129},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

}

func TestManyCorrId(t *testing.T) {
	t.Skip()
	t.Parallel()

	descriptor := &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{0, 1},
				{10, 11},
			},
		},
	}

	assert.Equal(t, []byte("0"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1024 + 100, 1024 + 104},
			},
		},
	}

	assert.Equal(t, []byte("1"), descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1028, 1129},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

	descriptor = &ExtendedBlockMessageDescriptor{
		Transactions: &Transactions{
			IndexRanges: []IndexRange{
				{100, 104},
				{1028, 1029},
			},
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

}
