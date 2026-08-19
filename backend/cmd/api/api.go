package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Ayushmangit/mirrormate.git/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxIdleConns int
	maxOpenConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/v1/swagger/doc.json"),
		))

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/users", app.registerUserHandler)
			r.Post("/login", app.loginUserHandler)

		})
	})
	return r

}

func (app *application) run(mux http.Handler) {
	srv := http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}
	fmt.Println("The server is listening on port :", srv.Addr)
	srv.ListenAndServe()
}
