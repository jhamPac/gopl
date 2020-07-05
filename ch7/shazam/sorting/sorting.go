package sorting

import (
	"github.com/jhampac/gopl/ch7/shazam/music"
)

// ByArtist sorts a slice of Track by artist
type ByArtist []*music.Track

func (x ByArtist) Len() int {
	return len(x)
}

func (x ByArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x ByArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
