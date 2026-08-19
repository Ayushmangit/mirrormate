package main

import (
	"fmt"
	"net/http"
)

func (app *application) UnAuthorizedBasicError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted",charset="UTF-8"`)
	WriteJsonError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) UnAuthorized(w http.ResponseWriter, r *http.Request, err error) {
	WriteJsonError(w, http.StatusUnauthorized, err.Error())
}

func (app *application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("%v", err.Error())
	WriteJsonError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	WriteJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) NotFound(w http.ResponseWriter, r *http.Request, err error) {
	WriteJsonError(w, http.StatusNotFound, "not found")
}

func (app *application) ForbiddenResponse(w http.ResponseWriter, r *http.Request) {
	WriteJsonError(w, http.StatusForbidden, "forbidden request")
}

func (app *application) RateLimitExceeded(w http.ResponseWriter, r *http.Request, retryAfter string) {
	w.Header().Set("Retry-After", retryAfter)
	WriteJsonError(w, http.StatusTooManyRequests, "rate limit exceeded")
}

func (app *application) Conflict(w http.ResponseWriter, r *http.Request, err error) {
	WriteJsonError(w, http.StatusConflict, err.Error())
}
