package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type UserKey string

const userCtx UserKey = "user"

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.UnAuthorized(w, r, errors.New("auth header missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.UnAuthorized(w, r, errors.New("missing token"))
			return
		}

		jwtToken, err := app.auth.ValidateToken(parts[1])
		if err != nil {
			app.UnAuthorized(w, r, err)
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			app.UnAuthorized(w, r, errors.New("invalid claims"))
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["sub"]), 10, 64)
		if err != nil {
			app.UnAuthorized(w, r, err)
			return
		}

		user, err := app.getUser(r.Context(), userID)
		if err != nil {
			app.UnAuthorized(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	user, err := app.store.Users.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, store.ErrNotFound
	}

	return user, nil
}
