package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Get a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path		int	true	"User ID"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
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
//	@Param			userID	path	int	true	"User ID"
//	@Param			request	body	UpdateUserPayload	true	"User update payload"
//	@Success		200		{object}	store.User
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/users/{userID} [patch]
func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	_, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			userID	path	int	true	"User ID"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
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

	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.InternalServerError(w, r, err)
		return
	}
}
