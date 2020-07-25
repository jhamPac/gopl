package arprint_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhampac/gopl/ch10/arprint"
	_ "github.com/jhampac/gopl/ch10/arprint/tar"
	_ "github.com/jhampac/gopl/ch10/arprint/zip"
)

func TestOpen(t *testing.T) {
	for _, archive := range []string{"rah.zip, rah.tar"} {
		b := &bytes.Buffer{}
		f, err := os.Open(filepath.Join("testdata", archive))
		if err != nil {
			t.Error(archive, err)
		}
		r, err := arprint.Open(f)
		if err != nil {
			t.Error(archive, err)
		}
		_, err = io.Copy(b, r)
		if err != nil {
			t.Error(archive, r)
		}
		want := `rah/b:
		contentsB
		rah/a:
		contentsA
		`

		got := b.String()
		if got != want {
			t.Errorf("%s: got %q, want %q", archive, got, want)
		}
	}
}
