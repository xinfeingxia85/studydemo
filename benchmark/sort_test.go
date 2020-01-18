package benchmark

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"sort"
	"testing"
)

func gernerateSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(1e9))
	}

	return s
}

func BenchmarkSort1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := gernerateSlice(1000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := gernerateSlice(10000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := gernerateSlice(100000)
		b.StartTimer()
		sort.Ints(s)
	}
}

func BenchmarkMathrand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Int63()
	}
}

func BenchmarkCrypttorand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := crand.Int(crand.Reader, big.NewInt(27))
		if err != nil {
			panic(err)
		}
	}
}
