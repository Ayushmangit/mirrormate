package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
	"github.com/go-chi/chi/v5"
)

type MessageResponse struct {
	Message string `json:"message"`
}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{userID} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}
	user, err := app.store.Users.GetByID(r.Context(), id)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update a user's information
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userID	path		int						true	"User ID"
//	@Param			request	body		store.UpdateUserPayload	true	"User update payload"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	error
//	@Failure		404		{object}	error
//	@Failure		409		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{userID} [patch]
func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	payload := store.UpdateUserPayload{}

	if err := ReadJson(w, r, &payload); err != nil {
		app.BadRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequest(w, r, err)
		return

	}

	if payload.Email == nil && payload.Username == nil {
		app.BadRequest(w, r, errors.New("at least one field is required"))
		return
	}

	user, err := app.store.Users.UpdateByID(
		r.Context(),
		id,
		payload,
	)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.NotFound(w, r, err)

		case errors.Is(err, store.ErrDuplicateEmail),
			errors.Is(err, store.ErrDuplicateUsername):
			app.Conflict(w, r, err)

		default:
			app.InternalServerError(w, r, err)
		}

		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path		int				true	"User ID"
//	@Success		200		{object}	MessageResponse	"user deleted successfully"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/{userID} [delete]
func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userID")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	if err := app.store.Users.DeleteByID(r.Context(), id); err != nil {
		app.InternalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, MessageResponse{
		Message: "user deleted successfully",
	}); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}

// ActivateUser godoc
//
//	@Summary		Activates a user account
//	@Description	Activates a user account using an invitation token
//	@Tags			users
//	@Produce		json
//	@Param			token	path		string			true	"Invitation Token"
//	@Success		200		{object}	MessageResponse	"user activated successfully"
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/users/activate/{token} [get]
func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	err := app.store.Users.Activate(r.Context(), token)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.NotFound(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, MessageResponse{
		Message: "user activated successfully",
	}); err != nil {
		app.InternalServerError(w, r, err)
	}
}
