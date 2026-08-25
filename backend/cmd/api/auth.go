package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
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

// LoginUser godoc
//
//	@Summary		Login a user
//	@Description	Authenticates a user and returns a success response
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginUserPayload	true	"User login payload"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/authentication/login [post]
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload LoginUserPayload
	if err := ReadJson(w, r, &payload); err != nil {
		app.BadRequest(w, r, errors.New("invalid request payload"))
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
			app.UnAuthorized(w, r, errors.New("invalid credentials"))
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
	if err := user.Password.Compare(payload.Password); err != nil {
		app.UnAuthorized(w, r, errors.New("invalid credentials"))
		return
	}

	if !user.IsActive {
		app.UnAuthorized(w, r, errors.New("user account is not activated"))
		return
	}
	//generate claims
	// generate JWT token with claims
	//send the user and jwt in response

	//TODO: create a LoginResponse type which will have the user and JWT

	if err := app.jsonResponse(w, http.StatusOK, map[string]string{
		"message": "login successful",
	}); err != nil {
		app.InternalServerError(w, r, err)
	}
}

type UserWithToken struct {
	*store.User `json:"user"`
	Token       string `json:"token"`
}

func generateRandomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

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
		app.BadRequest(w, r, errors.New("invalid request payload"))
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
	// 1. Generate cryptographically random plain token
	plainToken, err := generateRandomToken()
	if err != nil {
		app.InternalServerError(w, r, err)
		return
	}
	// 2. Hash plain token with SHA-256 before storing
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])
	// 3. Store user with invitation token (expiry: 3 days)

	//TODO: use the app.cfg.auth.token.expiry instead of hardcoding
	invitationExp := 3 * 24 * time.Hour

	if err := app.store.Users.CreateAndInvite(ctx, user, hashToken, invitationExp); err != nil {
		switch err {
		case store.ErrDuplicateEmail, store.ErrDuplicateUsername:
			app.BadRequest(w, r, err)
		default:
			app.InternalServerError(w, r, err)
		}
		return
	}
	// 4. Return response containing plain token
	response := UserWithToken{
		User:  user,
		Token: plainToken,
	}
	if err := app.jsonResponse(w, http.StatusCreated, response); err != nil {
		app.InternalServerError(w, r, err)
	}
}
