package benchmark

import (
	"hash"
	"math/rand"
	"testing"
)

func benchmarkHash(b *testing.B, h hash.Hash) {
	data := make([]byte, 1024)
	rand.Read(data)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		h.Write(data)
		h.Sum(nil)
	}
}
