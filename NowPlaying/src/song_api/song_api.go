package song_api

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//FindAll ...
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		songs, err2 := songModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, songs)

		}
	}
}

//Search ...
func Search(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		songs, err2 := songModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, songs)

		}
	}
}

//SearchYear ...
/*func SearchYear(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	smin := vars["min"]
	smax := vars["max"]
	min, _ := strconv.ParseInt(smin, 64, 64)
	max, _ := strconv.ParseInt(smax, 64, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		songs, err2 := songModel.SearchYear(min, max)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, songs)

		}
	}
}*/

//Create ...
func Create(response http.ResponseWriter, request *http.Request) {

	var song entities.Songs
	err := json.NewDecoder(request.Body).Decode(&song)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		err2 := songModel.Create(&song)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, song)

		}
	}
}

//Update ...
func Update(response http.ResponseWriter, request *http.Request) {

	var song entities.Songs
	err := json.NewDecoder(request.Body).Decode(&song)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		_, err2 := songModel.Update(&song)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, song)

		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, s interface{}) {
	response, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

//Delete ...
func Delete(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		songModel := models.SongModel{
			Db: db,
		}
		_, err2 := songModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)

		}
	}
}
