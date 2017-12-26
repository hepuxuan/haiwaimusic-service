package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"haiwaimusic-service/models"
	"net/http"
	"fmt"
)

type PlayListController struct {
	Db *sql.DB
}

func (uc PlayListController) FindPlayList(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	userId := pathParams["userId"]

	playList, err := models.FindPlayListByUserId(userId, uc.Db)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(playList)
}

func (uc PlayListController) AddToPlayList(w http.ResponseWriter, r *http.Request)  {
	pathParams := mux.Vars(r)
	userId := pathParams["userId"]

	var song models.Song

	json.NewDecoder(r.Body).Decode(&song)

	song.UserId = userId

	err := models.AddPlayList(song, uc.Db)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func (uc PlayListController) RemoveFromPlayList(w http.ResponseWriter, r *http.Request)  {
	pathParams := mux.Vars(r)
	userId := pathParams["userId"]
	songId := pathParams["songId"]

	song := models.Song{}

	json.NewDecoder(r.Body).Decode(&song)

	song.UserId = userId

	err := models.RemovePlayList(songId, userId, uc.Db)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
