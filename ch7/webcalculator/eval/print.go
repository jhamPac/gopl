package eval

import "bytes"

func Format(e Expr) string {
	var buf bytes.Buffer
	write(&buf, e)
	return buf.String()
}
