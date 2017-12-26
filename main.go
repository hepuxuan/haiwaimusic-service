package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"haiwaimusic-service/controllers"
	"net/http"
)

import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/haiwaimusic")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()

	userController := controllers.UserController{
		Db: db,
	}

	router.HandleFunc("/users/{userId}", userController.FindUser).Methods("GET")
	router.HandleFunc("/users", userController.CreateOrUpdate).Methods("POST")

	playListController := controllers.PlayListController{
		Db: db,
	}

	router.HandleFunc("/users/{userId}/playList", playListController.FindPlayList).Methods("GET")
	router.HandleFunc("/users/{userId}/playList", playListController.AddToPlayList).Methods("POST")
	router.HandleFunc("/users/{userId}/playList/{songId}", playListController.RemoveFromPlayList).Methods("DELETE")
	http.ListenAndServe(":8081", router)
}
