package controllers

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"haiwaimusic-service/models"
	"fmt"
	"encoding/json"
)

type UserController struct {
	Db *sql.DB
}

func (uc UserController) FindUser(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	userId := pathParams["userId"]

	user := models.FindUserByUuid(userId, uc.Db)

	json.NewEncoder(w).Encode(user)
}

func (uc UserController) CreateOrUpdate(w http.ResponseWriter, r *http.Request)  {
	user := models.User{}

	json.NewDecoder(r.Body).Decode(&user)

	err := models.CreateOrUpdate(user, uc.Db)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
