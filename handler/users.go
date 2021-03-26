package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PicPay/software-engineer-challenge/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var UserIDKey = "UserID"

func users(router chi.Router) {
	router.Get("/", getUsers)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var p models.Req
	var err error
	resOrdered, resAux, res := &models.UserList{},
		&models.UserList{},
		&models.UserList{}

	p.Page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if p.Page != 0 {
		var start int
		MAX_PAGE_SIZE := 538545
		page := 15
		if p.Page == 1 {
			start = 0
		} else if p.Page > MAX_PAGE_SIZE {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("page size exceed the max size for the table -> %d", MAX_PAGE_SIZE)))
			return
		} else {
			start = 15 * (p.Page - 1)
		}
		log.Println(start, p.Page)
		User, err := dbInstance.GetAllUsersLimit(&models.User{}, start, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Users = User.Users
	} else {
		render.Render(w, r, ErrorRenderer(errors.New("page need to be at least 1")))
		return
	}

	log.Printf("%d users found on db -> %v", len(res.Users), res.Users)
	resp := Tfile()

	Parsing(&resp, res, resOrdered, resAux)

	resOrdered.Users = append(resOrdered.Users, resAux.Users...)
	if err := render.Render(w, r, resOrdered); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
