package sorting

import (
	"github.com/jhampac/gopl/ch7/shazam/music"
)

type byArtist []*music.Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
