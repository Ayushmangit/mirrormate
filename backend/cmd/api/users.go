package main

import (
	"net/http"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: Add validation later on

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user account
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateUserPayload	true	"User registration payload"
//	@Success		201		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		409		{object}	error
//	@Failure		500		{object}	error
//	@Router			/users [post]
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload CreateUserPayload
	if err := ReadJson(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Role: store.Role{
			Name: "user",
		},
	}

	if err := app.store.Users.Create(ctx, user); err != nil {
		//NOTE: Never send errors like email already exists , intruder can then start guessing the passwords
		switch err {
		case store.ErrDuplicateEmail:
			// http.Error(w, "email already exists", http.StatusConflict)
			app.BadRequest(w, r, err)

		case store.ErrDuplicateUsername:
			app.BadRequest(w, r, err)

		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
}
