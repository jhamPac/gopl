package limit

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "🌈🗿x🐲🍔eger"
	b := &bytes.Buffer{}
	r := NewLimitReader(strings.NewReader(s), 9)
	n, _ := b.ReadFrom(r)
	if n != 9 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	if b.String() != "🌈🗿x" {
		t.Log(b.String())
		t.Fatalf("buffer string is incorrect")
	}
}
