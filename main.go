package main

import (
	"log"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/create", models.Create).Methods("POST")
	router.HandleFunc("/findall", models.FindAll).Methods("GET")
	router.HandleFunc("/update/{songid}", models.Update).Methods("PUT")
	router.HandleFunc("/deleteone/{songid}", models.Delete).Methods("DELETE")
	router.HandleFunc("/deleteall", models.DeleteAll).Methods("DELETE")
	router.HandleFunc("/count", models.Count).Methods("GET")

	log.Fatal(http.ListenAndServe(":3002", router))

}
