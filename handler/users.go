package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PicPay/software-engineer-challenge/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var UserIDKey = "UserID"

func users(router chi.Router) {
	router.Get("/", getAllItems)
	router.Post("/", createItem)
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserId := chi.URLParam(r, "UserId")
		if UserId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("User ID is required")))
			return
		}
		id, err := strconv.Atoi(UserId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid User ID")))
		}
		ctx := context.WithValue(r.Context(), UserIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllUsers()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, items); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createItem(w http.ResponseWriter, r *http.Request) {
	item := &models.User{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddUser(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
