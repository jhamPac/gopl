package counter

import "testing"

func TestLineCounter(t *testing.T) {
	c := &LineCounter{}
	p := []byte("one\ntwo\nthree\n")
	n, err := c.Write(p)
	if n != len(p) {
		t.Logf("len: %d != %d", n, len(p))
		t.Fail()
	}
	if err != nil {
		t.Log("err: ", err)
		t.Fail()
	}
	if c.N() != 3 {
		t.Logf("lines: %d != 3", c.N())
	}
}
