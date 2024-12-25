package main

import (
	"github.com/adriannbhp/social-apps/internal/db"
	"github.com/adriannbhp/social-apps/internal/env"
	"github.com/adriannbhp/social-apps/internal/store"
	"log"
)

const version = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:root@localhost/social-network?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxLifeTime:  env.GetString("DB_MAX_LIFE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	log.Println("Starting database connection...")
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxLifeTime,
	)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database connection established.")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
