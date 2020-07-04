package reader

import (
	"bytes"
	"testing"
)

func TestNewReader(t *testing.T) {
	s := "hi there"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Logf("n=%d err=%s", n, err)
		t.Fail()
	}
	if b.String() != s {
		t.Logf("%q != %q", b.String(), s)
	}
}
