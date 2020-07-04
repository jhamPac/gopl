package limit

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hi there"
	b := &bytes.Buffer{}
	r := NewLimitReader(strings.NewReader(s), 4)
	n, _ := b.ReadFrom(r)
	if n != 4 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	if b.String() != "hi t" {
		t.Fatalf("buffer string is incorrect")
	}
}
