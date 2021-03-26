package handler

import (
	"net/http"
	"sort"

	"github.com/PicPay/software-engineer-challenge/db"
	"github.com/PicPay/software-engineer-challenge/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/users", users)
	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}

func Parsing(resp *chan []string, res, resOrdered, resAux *models.UserList) {
	ids_1 := map[string]int{}
	aux, i := 0, 0
	for r := range *resp {
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
}
