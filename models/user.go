package models

import "database/sql"

type User struct {
	Id int32
	Uuid string
	Name string
	ImageUrl string
	Email string
}

func FindUserByUuid(uuid string, db *sql.DB) User {
	user := User{}

	row := db.QueryRow("SELECT id, uuid, name, image_url, email FROM user WHERE uuid=?", uuid)

	row.Scan(&user.Id, &user.Uuid, &user.Name, &user.ImageUrl, &user.Email)

	return user
}

func CreateOrUpdate(user User, db *sql.DB) error  {
	var err error
	existingUser := FindUserByUuid(user.Uuid, db)

	if existingUser.Id == 0 {
		_, err = db.Exec("INSERT INTO user (uuid, name, image_url, email) VALUES(?, ?, ?, ?)",
			user.Uuid, user.Name, user.ImageUrl, user.Email)
	} else {
		_, err = db.Exec("UPDATE user SET uuid = ?, name = ?, image_url = ?, email = ? WHERE id = ?",
			user.Uuid, user.Name, user.ImageUrl, user.Email, user.Id)
	}

	return err
}
