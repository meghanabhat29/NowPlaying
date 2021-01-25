package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Song struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	year   int    `json:"year"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Intern@paytm29@tcp(127.0.0.1:3306)/Playlist")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/songs", getPosts).Methods("GET")
	router.HandleFunc("/songs", createPost).Methods("POST")
	router.HandleFunc("/songs/{id}", getPost).Methods("GET")
	router.HandleFunc("/songs/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/songs/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var songs []Song
	result, err := db.Query("SELECT id, name, artist, album, genre, year from Playlist.songs")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var song Song
		err := result.Scan(&song.Id, &song.Name, &song.Artist, &song.Album, &song.Genre, &song.year)
		if err != nil {
			panic(err.Error())
		}
		songs = append(songs, song)
	}
	json.NewEncoder(w).Encode(songs)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO Playlist.songs(name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	_, err = stmt.Exec(name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New song added")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT Id, Name FROM Playlist.songs WHERE Id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var song Song
	for result.Next() {
		err := result.Scan(&song.Id, &song.Name)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(song)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE Playlist.songs SET Name = ? WHERE Id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTitle := keyVal["name"]
	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM Playlist.songs WHERE Id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
