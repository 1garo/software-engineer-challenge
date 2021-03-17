package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/PicPay/software-engineer-challenge/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var UserIDKey = "UserID"

func users(router chi.Router) {
	router.Get("/", getAllUsers)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
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
	/*
		if the id sent on the request do not exist on the db,
		it just skip it and go for the next one that can be found
	*/
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

	ids_1 := map[string]int{}
	aux, i := 0, 0
	for r := range resp {
		ids_1[r[0]] = aux
		aux += 1
	}

	sort.Slice(
		res.Users,
		func(i, j int) bool { return ids_1[res.Users[i].ID] < ids_1[res.Users[j].ID] })

	for _, user := range res.Users {
		if _, ok := ids_1[user.ID]; ok {
			resOrdered.Users = append(resOrdered.Users, user)
			i += 1
		} else {
			resAux.Users = append(resAux.Users, user)
		}
	}

	resOrdered.Users = append(resOrdered.Users, resAux.Users...)
	if err := render.Render(w, r, resOrdered); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
