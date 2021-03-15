package models

import (
	"fmt"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserList struct {
	Users []User `json:"Users"`
}

func (i *User) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}
func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
