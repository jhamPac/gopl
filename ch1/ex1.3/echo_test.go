package echo

import (
	"testing"
)

var (
	args = []string{"exec arg0 arg1 arg2 arg3"}
)

func TestEcho1(t *testing.T) {
	if len(args) <= 0 {
		t.Error("not enough args")
	}
	echo1(args)
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
