package popcount

import (
	"testing"
)

var n uint64 = 0x1234567890abcdef

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcount(n)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcount2(n)
	}
}
