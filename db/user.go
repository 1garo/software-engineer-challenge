package db

import (
	// "database/sql"
	"strconv"
	"strings"

	"github.com/PicPay/software-engineer-challenge/models"
)

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

func (db Database) AddUserInDb(records chan []string) error {
	var err error
	sqlStr := "INSERT INTO Users (id, name, username) VALUES "
	vals := []interface{}{}

	for record := range records {
		sqlStr += "(?, ?, ?),"
		vals = append(vals, record[0], record[1], record[2])
	}

	//trim the last ,
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	//Replacing ? with $n for postgres
	sqlStr = replaceSQL(sqlStr, "?")

	//format all vals at once
	stmt, err := db.Conn.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}
	return nil
}

// ReplaceSQL replaces the instance occurrence of any string pattern with an increasing $n based sequence
func replaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
