package main

import "github.com/jhampac/gopl/ch7/shazam/music"

type clickedSort struct {
	tracks []*music.Track
	less   func(track1, track2 *music.Track) bool
}

func (x clickedSort) Len() int {
	return len(x.tracks)
}

func (x clickedSort) Less(i, j int) bool {
	return x.less(x.tracks[i], x.tracks[j])
}

func (x clickedSort) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}
