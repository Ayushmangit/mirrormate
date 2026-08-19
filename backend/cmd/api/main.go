package main

import (
	"github.com/Ayushmangit/mirrormate.git/internal/db"
	"github.com/Ayushmangit/mirrormate.git/internal/env"
	"github.com/Ayushmangit/mirrormate.git/internal/store"

	_ "github.com/Ayushmangit/mirrormate.git/docs"
)

const version = "0.0.1" // semver

//	@title			MirrorMate
//	@description	SocialNetwork for hiring
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable"),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	app := application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	app.run(mux)
}
