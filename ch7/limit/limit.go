package limit

import "io"

type limitReader struct {
	r     io.Reader
	n     int
	limit int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p[:l.limit])
	l.n += n
	if l.n >= l.limit {
		err = io.EOF
	}
	return
}

// NewLimitReader returns a new limitReader type
func NewLimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
