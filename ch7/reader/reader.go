package reader

import "io"

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

// NewReader creates and returns a new type of stringReader
func NewReader(s string) io.Reader {
	return &stringReader{s}
}
