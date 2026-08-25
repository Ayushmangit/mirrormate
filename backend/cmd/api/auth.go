package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type LoginResponse struct {
	*store.User `json:"user"`
	AccessToken string `json:"access_token"`
}

// LoginUser godoc
//
//	@Summary		Login a user
//	@Description	Authenticates a user and returns a success response
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginUserPayload	true	"User login payload"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/login [post]
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload LoginUserPayload
	if err := ReadJson(w, r, &payload); err != nil {
		app.BadRequest(w, r, store.ErrBadRequest)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.BadRequest(w, r, err)
		return
	}

	user, err := app.store.Users.GetByEmail(ctx, payload.Email)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.BadRequest(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
	if err := user.Password.Compare(payload.Password); err != nil {
		app.UnAuthorized(w, r, store.ErrUnAuthorized)
		return
	}

	if !user.IsActive {
		app.UnAuthorized(w, r, store.ErrNotActivated)
		return
	}
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(app.config.auth.token.exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": app.config.auth.token.iss,
		"aud": app.config.auth.token.aud,
	}
	accessToken, err := app.auth.GenerateToken(claims)
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	response := LoginResponse{
		user,
		accessToken,
	}

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.InternalServerError(w, r, err)
	}
}

type RegisterResponse struct {
	*store.User `json:"user"`
	Token       string `json:"token"`
}

// RegisterUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user account
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User registration payload"
//	@Success		201		{object}	RegisterResponse
//	@Failure		400		{object}	error
//	@Failure		409		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/users [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload RegisterUserPayload
	if err := ReadJson(w, r, &payload); err != nil {
		app.BadRequest(w, r, store.ErrBadRequest)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.BadRequest(w, r, err)
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
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(ctx,
		user,
		hashToken,
		app.config.auth.token.exp); err != nil {
		switch err {
		case store.ErrDuplicateEmail, store.ErrDuplicateUsername:
			app.BadRequest(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
	response := RegisterResponse{
		User:  user,
		Token: plainToken,
	}
	if err := app.jsonResponse(w, http.StatusCreated, response); err != nil {
		app.InternalServerError(w, r, err)
	}
}
