package benchmark

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	data := make([]byte, 1024)
	_, _ = rand.Read(data)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		base64.StdEncoding.EncodeToString(data)
	}
}

func BenchmarkDecode(b *testing.B) {
	data := make([]byte, 1024)
	_, _ = rand.Read(data)
	encoded := base64.StdEncoding.EncodeToString(data)

	for n := 0; n < b.N; n++ {
		_, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			panic(err)
		}
	}
}
