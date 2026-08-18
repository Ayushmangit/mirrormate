package main

import (
	"encoding/json"
	"net/http"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
	}

	ctx := r.Context()
	if err := app.store.Users.Create(ctx, user); err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			http.Error(w, "email already exixts", http.StatusConflict)

		case store.ErrDuplicateUsername:
			http.Error(w, "username already exists", http.StatusConflict)

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

}
