package benchmark

import (
	"fmt"
	"testing"
)

func benchmarkCalculate(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}

func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	resulst := Calculate(2)
	if expected != resulst {
		t.Error("Failed")
	}
}

func BenchmarkCalculate100(b *testing.B) {
	benchmarkCalculate(100, b)
}

func BenchmarkCalculateNagative100(b *testing.B) {
	benchmarkCalculate(-100, b)
}

func BenchmarkCalculateNagative1(b *testing.B) {
	benchmarkCalculate(-1, b)
}
