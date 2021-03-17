package db

import (
	"fmt"

	"github.com/PicPay/software-engineer-challenge/models"
)

func (db Database) GetAllUsersLimit(User *models.User, start, page int) (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM users order by id offset $1 limit $2", start, page)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var User models.User
		err := rows.Scan(&User.ID, &User.Name, &User.Username)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, User)
	}
	return list, nil
}

func (db Database) GetAllUsers(User *models.User) (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM Users where id = $1", User.ID)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var User models.User
		err := rows.Scan(&User.ID, &User.Name, &User.Username)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, User)
	}
	return list, nil
}

func (db Database) CopyTable() error {
	_, err := db.Conn.Exec(fmt.Sprintf("COPY users(id, username, name) FROM '%s' DELIMITER ',' CSV", "/tmp/users.csv"))
	if err != nil {
		return err
	}
	return nil
}
