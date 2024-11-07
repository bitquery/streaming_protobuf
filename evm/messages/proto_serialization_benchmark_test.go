package evm_messages

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/pierrec/lz4/v4"
	"io"
	"os"
)

func readMessageFromFile(tb testing.TB) proto.Message {
	// filename := "/tmp/dev.eth.transactions/000020876000/000020876330_0x24d6a679a658cb4130c6a8b580bfc0d19afc55ea55510a1a862dc9a1082fdc07_bea89c5e2d01575a7cb7334fe2b629555686a02db218014f1a140e8c2eab5ab6.block.lz4"
	filename := "/tmp/000020879999_0x54b08eab4e16cad5801b029e25023a208c5458db634ab60c26333154452ed895_db6107eb30cf08d0b9a5c99e6c162b1aaf9d7d370f11456cd1aa94ff40868cbc.block.lz4"

	reader, err := os.Open(filename)
	if err != nil {
		tb.Fatalf("Failed to open file: %v", err)
	}
	defer reader.Close()

	lz4Reader := lz4.NewReader(reader)
	body, err := io.ReadAll(lz4Reader)
	if err != nil {
		tb.Fatalf("Failed to read file: %v", err)
	}

	// message := &ParsedAbiBlockMessage{}
	message := &BlockMessage{}
	err = proto.Unmarshal(body, message)
	if err != nil {
		tb.Fatalf("Failed to unmarshal message: %v", err)
	}

	return message
}

// func TestReadMessageFromFile(t *testing.T) {
// 	message := readMessageFromFile(t)
// 	t.Logf("Message: %v", message)
// }

// func TestEqualSerialization(t *testing.T) {
// 	message := readMessageFromFile(t)
//
// 	standard := StandardSerialization(t, message)
// 	deterministic := DeterministicSerialization(t, message)
//
// 	standardMessage := &BlockMessage{}
// 	err := proto.Unmarshal(standard, standardMessage)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal message: %v", err)
// 	}
//
// 	deterministicMessage := &BlockMessage{}
// 	err = proto.Unmarshal(deterministic, deterministicMessage)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal message: %v", err)
// 	}
//
// 	if !proto.Equal(standardMessage, deterministicMessage) {
// 		t.Fatalf("Standard and deterministic serialization are not equal")
// 	}
// }

func BenchmarkStandardSerialization(b *testing.B) {
	message := readMessageFromFile(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		StandardSerialization(b, message)
	}
}

func BenchmarkDeterministicSerialization(b *testing.B) {
	message := readMessageFromFile(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		DeterministicSerialization(b, message)
	}
}

func StandardSerialization(tb testing.TB, message proto.Message) []byte {
	tb.Helper()

	res, err := proto.Marshal(message)
	if err != nil {
		tb.Fatalf("Failed to marshal message: %v", err)
	}

	return res
}

func DeterministicSerialization(tb testing.TB, message proto.Message) []byte {
	tb.Helper()

	buf := proto.NewBuffer(nil)
	buf.SetDeterministic(true)

	err := buf.Marshal(message)
	if err != nil {
		tb.Fatalf("Failed to marshal message: %v", err)
	}

	return buf.Bytes()
}
