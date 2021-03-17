package models

import (
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

type Req struct {
	Page int
}

func (i *User) Bind(r *http.Request) error {
	return nil
}
func (*UserList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (Req) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (i *Req) Bind(r *http.Request) error {
	return nil
}
