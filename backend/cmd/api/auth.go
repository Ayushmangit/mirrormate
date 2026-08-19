package main

import (
	"net/http"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
)

type RegisterUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: Add validation later on

// RegisterUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user account
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User registration payload"
//	@Success		201		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		409		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/users [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload RegisterUserPayload
	if err := ReadJson(w, r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		Role: store.Role{
			Name: "user",
		},
	}
	if err := user.Password.Set(payload.Password); err != nil {
		app.InternalServerError(w, r, err)
		return
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

	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.InternalServerError(w, r, err)
	}
}
