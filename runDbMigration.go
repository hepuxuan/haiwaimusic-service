package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"haiwaimusic-service/migrations"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost)/haiwaimusic")
	if err != nil {
		panic(err)
	}

	fmt.Println("running migrations")

	migrations.CreateUser(db)
	migrations.CreatePlayList(db)
}
