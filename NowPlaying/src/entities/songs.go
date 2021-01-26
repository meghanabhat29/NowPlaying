package entities

import (
	"fmt"
)

//Songs ...
type Songs struct {
	ID     int64  `json: "id"`
	Name   string `json: "name"`
	Artist string `json: "artist"`
	Album  string `json: "album"`
	Genre  string `json: "genre"`
	Year   int64  `json: "year"`
}

//ToString .. ...
func (song Songs) ToString() string {
	return fmt.Sprintf("Id: %d\nName: %s\nArtist: %s\nAlbum: %s\nGenre: %s\nYear: %d\n",
		song.ID, song.Name, song.Artist, song.Album, song.Genre, song.Year)
}
