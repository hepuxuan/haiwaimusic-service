package migrations

import (
	"database/sql"
	"fmt"
)

func CreatePlayList(db *sql.DB) {
	fmt.Println("creating playList table")

	_, err := db.Exec("CREATE TABLE play_list (" +
		"id INT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
		"user_id VARCHAR(255)," +
		"song VARCHAR(255)," +
		"singer VARCHAR(255)," +
		"mid VARCHAR(255)," +
		"song_id BIGINT," +
		"image_id BIGINT," +
		"INDEX (user_id)" +
		//"FOREIGN KEY (user_id) REFERENCES user(uuid) ON DELETE CASCADE" +
		")")

	if err != nil {
		panic(err)
	}
}
