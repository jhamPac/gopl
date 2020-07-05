package main

import (
	"sort"

	"github.com/jhampac/gopl/ch7/shazam/music"
	"github.com/jhampac/gopl/ch7/shazam/sorting"
)

func main() {
	tracks := []*music.Track{
		music.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
		music.NewTrack("Go", "Moby", "Moby Play", 1992, "3m10s"),
		music.NewTrack("Go Ahead", "Alicia Keys", "As Iam", 2007, "4m36s"),
		music.NewTrack("Read 2 Go", "Martin Solveig", "Samsh", 2011, "4m00s"),
	}
	sort.Sort(sorting.ByArtist(tracks))
	music.PrintTracks(tracks)
}
