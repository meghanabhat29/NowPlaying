package main

import (
	"fmt"
	"net/http"
	"song_api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/song/findall", song_api.FindAll).Methods("GET")
	router.HandleFunc("/api/song/search/{keyword}", song_api.Search).Methods("GET")
	//router.HandleFunc("/api/song/search/{min}/{max}", song_api.SearchYear).Methods("GET")
	router.HandleFunc("/api/song/create", song_api.Create).Methods("POST")
	router.HandleFunc("/api/song/update", song_api.Update).Methods("PUT")
	router.HandleFunc("/api/song/delete/{id}", song_api.Delete).Methods("DELETE")
	err := http.ListenAndServe(":3005", router)
	if err != nil {
		fmt.Println(err)
	}
}
