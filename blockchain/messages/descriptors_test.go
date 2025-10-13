package blockchain_messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneCorrId(t *testing.T) {
	t.Parallel()

	descriptor := &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901234",
		},
	}

	assert.Equal(t, []byte("red_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901235",
		},
	}

	assert.Equal(t, []byte("blue_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901236",
		},
	}

	assert.Equal(t, []byte("green_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901237",
		},
	}

	assert.Equal(t, []byte("yellow_stream"), descriptor.CorrelationId())
}

func TestManyCorrId(t *testing.T) {
	t.Parallel()

	descriptor := &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901234",
			"0x1234567890123456789012345678901234567890123456789012345678901238",
			"0x123456789012345678901234567890123456789012345678901234567890123c",
		},
	}

	assert.Equal(t, []byte("red_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1234567890123456789012345678901234567890123456789012345678901235",
			"0x1234567890123456789012345678901234567890123456789012345678901234",
		},
	}

	assert.Nil(t, descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0xb432db11c6c6460f6bbf951a71b3ea3df9e0e414c1a4616728795111a7518f30",
			"0x74d35692f26ae3b6f0d276091a73093205b31ed320f127d5ce4c6454a250d98c",
			"0x8140ba1c7ea8104437e7c576fc43cb05d75276dcff3d99eec83ba7c4784787a8",
			"0x74a96c3df0649f9e3de102da99f2e6cac350c1537e5643e6c20b5b8b57dbf73c",
		},
	}

	assert.Equal(t, []byte("red_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0xd051e1a031d589ad246518525b792518984301ca5e833bbf64719783b20cc662",
			"0xad4f25e8a958dfd99baabf1a2479dda2765ecb05feb17efe2996331dc3eb5f4a",
		},
	}

	assert.Equal(t, []byte("green_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x1ed6796ba1bfae3e6c3ea00c077e3fb7e24522c6f7d332148ed248f15cd748f9",
			"0x9724e6f6835be2aad985bfe51a0b8498b20b49f71f1001394005913dc7a371a1",
		},
	}

	assert.Equal(t, []byte("blue_stream"), descriptor.CorrelationId())

	descriptor = &BlockMessageDescriptor{
		TransactionHashes: []string{
			"0x97026efd8249efac39bd45750b40baeb3a2def0357ceb1779797edf4802bd8be",
			"0xd7bebfb0968ad3f385dc2971ed9314bda30fa063a3969e14a5ca7a884bdd4402",
		},
	}

	assert.Equal(t, []byte("green_stream"), descriptor.CorrelationId())
}
