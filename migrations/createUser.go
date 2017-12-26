package migrations

import (
	"database/sql"
	"fmt"
)

func CreateUser(db *sql.DB) {
	fmt.Println("creating user table")

	_, err := db.Exec("CREATE TABLE user (" +
		"id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"uuid VARCHAR(255) NOT NULL," +
		"name VARCHAR(255)," +
		"image_url VARCHAR(255)," +
		"email VARCHAR(255)," +
		"INDEX (uuid)" +
		")")

	if err != nil {
		panic(err)
	}
}

