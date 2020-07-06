package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/jhampac/gopl/ch7/shazam/music"
)

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

var tracks = []*music.Track{
	music.NewTrack("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
	music.NewTrack("Go", "Moby", "Moby Play", 1992, "3m10s"),
	music.NewTrack("Go Ahead", "Alicia Keys", "As Iam", 2007, "4m36s"),
	music.NewTrack("Read 2 Go", "Martin Solveig", "Samsh", 2011, "4m00s"),
}

var trackTable = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style media="screen" type="text/css">
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Tracks</h1>
		<table>
		  <thead>
				<tr>
					<th><a href="/?sort=title">Title</a></th>
					<th><a href="/?sort=artist">Artist</a></th>
					<th><a href="/?sort=album">Album</a></th>
					<th><a href="/?sort=year">Year</a></th>
					<th><a href="/?sort=length">Length</a></th>
				</tr>
			</thead>
			<tbody>
				{{range .}}
				<tr>
					<td>{{.Title}}</td>
					<td>{{.Artist}}</td>
					<td>{{.Album}}</td>
					<td>{{.Year}}</td>
					<td>{{.Length}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Shazam tracks list is running!")
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sortBy := r.FormValue("sort")

	sort.Sort(clickedSort{tracks, func(track1, track2 *music.Track) bool {
		switch sortBy {
		case "title":
			return track1.Title < track2.Title
		case "year":
			return track1.Year < track2.Year
		case "length":
			return track1.Length < track2.Length
		case "artist":
			return track1.Artist < track2.Artist
		case "album":
			return track1.Album < track2.Album
		}
		return false
	}})
	printTracks(w, tracks)
}

func printTracks(w io.Writer, tracks []*music.Track) {
	if err := trackTable.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}
