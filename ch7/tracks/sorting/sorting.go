package sorting

import "time"

// Track represents a music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}
