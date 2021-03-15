package db

import (
	// "database/sql"
	"database/sql"

	"github.com/PicPay/software-engineer-challenge/models"
)

// func (db Database) GetAllUsers() (*models.UserList, error) {
// 	list := &models.UserList{}
// 	rows, err := db.Conn.Query("SELECT * FROM Users ORDER BY ID DESC")
// 	if err != nil {
// 		return list, err
// 	}
// 	for rows.Next() {
// 		var User models.User
// 		err := rows.Scan(&User.Id, &User.Name, &User.Username)
// 		if err != nil {
// 			return list, err
// 		}
// 		list.Users = append(list.Users, User)
// 	}
// 	return list, nil
// }
// func (db Database) AddUser(User *models.User) error {
// 	var id string
// 	query := `INSERT INTO Users (id, name, username) VALUES ($1, $2, $3) RETURNING id`
// 	err := db.Conn.QueryRow(query, User.Id, User.Name, User.Username).Scan(&id)
// 	if err != nil {
// 		return err
// 	}
// 	User.Id = id
// 	return nil
// }
func (db Database) GetAllUsers() (*models.UserList, error) {
	list := &models.UserList{}
	rows, err := db.Conn.Query("SELECT * FROM Users ORDER BY ID DESC")
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
func (db Database) AddUser(User *models.User) error {
	var id string
	query := `INSERT INTO Users (id, name, username) VALUES ($1, $2, $3) RETURNING id`
	err := db.Conn.QueryRow(query, User.ID, User.Name, User.Username).Scan(&id)
	if err != nil {
		return err
	}
	User.ID = id
	return nil
}
func (db Database) GetUserById(UserId int) (models.User, error) {
	User := models.User{}
	query := `SELECT * FROM Users WHERE id = $1;`
	row := db.Conn.QueryRow(query, UserId)
	switch err := row.Scan(&User.ID, &User.Name, &User.Username); err {
	case sql.ErrNoRows:
		return User, ErrNoMatch
	default:
		return User, err
	}
}
func (db Database) DeleteUser(UserId int) error {
	query := `DELETE FROM Users WHERE id = $1;`
	_, err := db.Conn.Exec(query, UserId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
func (db Database) UpdateUser(UserId int, UserData models.User) (models.User, error) {
	User := models.User{}
	query := `UPDATE Users SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, UserData.Name, UserData.Username, UserId).Scan(&User.ID, &User.Name, &User.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return User, ErrNoMatch
		}
		return User, err
	}
	return User, nil
}
