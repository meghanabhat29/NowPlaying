package models

import (
	"encoding/json"
	"entities"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

var session *gocql.Session

//Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	var song entities.Songs
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &song)
	if err := session.Query("INSERT INTO songs(name, artist, album, genre, year) VALUES(?, ?, ?, ?,?)",
		song.Name, song.Artist, song.Album, song.Genre, song.Year).Exec(); err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(song, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

//FindAll ...
func FindAll(w http.ResponseWriter, r *http.Request) {
	var songs []entities.Songs
	m := map[string]interface{}{}

	iter := session.Query("SELECT * FROM songs").Iter()
	for iter.MapScan(m) {
		songs = append(songs, entities.Songs{
			SongID: m["songid"].(int),
			Name:   m["name"].(string),
			Artist: m["artist"].(string),
			Album:  m["album"].(string),
			Genre:  m["genre"].(string),
			Year:   m["year"].(string),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(songs, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

//Count ...
func Count(w http.ResponseWriter, r *http.Request) {

	var Count string
	err := session.Query("SELECT count(*) FROM songs").Scan(&Count)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s ", Count)
}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	SongID := mux.Vars(r)["songid"]
	if err := session.Query("DELETE FROM songs WHERE songid = ?", SongID).Exec(); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Deleted %s ", SongID)
}

//DeleteAll ...
func DeleteAll(w http.ResponseWriter, r *http.Request) {

	if err := session.Query("TRUNCATE songs").Exec(); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "All the songs are deleted :(")

}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	SongID := mux.Vars(r)["songid"]
	var updatesong entities.Songs
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &updatesong)
	if err := session.Query("UPDATE songs SET name = ?, artist = ?, album = ? , genre = ?, year = ? WHERE songid = ?",
		updatesong.Name, updatesong.Artist, updatesong.Album, updatesong.Genre, updatesong.Year, SongID).Exec(); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Updated songs")

}
