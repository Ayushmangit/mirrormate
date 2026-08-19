package main

import (
	"net/http"
)

// healthCheck godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.InternalServerError(w, r, err)
	}
}
