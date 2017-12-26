package models

import "database/sql"

type Song struct {
	Song    string `json:"song"`
	Singer  string `json:"singer"`
	Mid     string `json:"mid"`
	SongId  int64 `json:"songId"`
	ImageId int64 `json:"imageId"`
	UserId  string `json:"userId"`
}

type PlayList struct {
	Songs []Song `json:"songs"`
}

func FindPlayListByUserId(id string, db *sql.DB) (PlayList, error) {
	playList := PlayList{}

	rows, err := db.Query("SELECT song, singer, mid, song_id, image_id, user_id FROM play_list WHERE user_id=?", id)
	if err != nil {
		return playList, err
	}

	defer rows.Close()

	songs := make([]Song, 0)

	for rows.Next() {
		song := Song{}
		rows.Scan(&song.Song, &song.Singer, &song.Mid, &song.SongId, &song.ImageId, &song.UserId)
		songs = append(songs, song)
	}

	playList.Songs = songs

	return playList, nil
}

func AddPlayList(song Song, db *sql.DB) error  {
	_, err := db.Exec("INSERT INTO play_list (user_id, song, singer, mid, song_id, image_id) VALUES(?, ?, ?, ?, ?, ?)",
		song.UserId, song.Song, song.Singer, song.Mid, song.SongId, song.ImageId)
	return err
}

func RemovePlayList(songId string, userId string, db *sql.DB) error  {
	_, err := db.Exec("DELETE FROM play_list WHERE user_id=? AND mid=?", userId, songId)
	return err
}
