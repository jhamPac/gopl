package track

import (
	"time"
)

// Track represents a music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// New creates and returns a pointer to a Track type
func New(title, artist, album string, year int, t string) *Track {
	return &Track{
		Title:  title,
		Artist: artist,
		Album:  album,
		Year:   year,
		Length: length(t),
	}
}
