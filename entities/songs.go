package entities

//Songs ...
type Songs struct {
	SongID int    `json:"song_id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}
