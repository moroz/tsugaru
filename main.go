package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oauth-provider/config"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

func logQuery(_ context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	if msg == "Query" {
		log.Printf("[%s] %s %v", level, data["sql"], data["args"])
		return
	}
	log.Printf("[%s] %s: %v", level, msg, data)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	cfg, err := pgxpool.ParseConfig(config.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	tracer := tracelog.TraceLog{}
	tracer.Logger = tracelog.LoggerFunc(logQuery)
	tracer.LogLevel = tracelog.LogLevelInfo
	cfg.ConnConfig.Tracer = &tracer

	conn, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec(context.Background(), "select now()")

	r := chi.NewRouter()
	r.Get("/", handleGet)
	log.Printf("Listening on %s", config.LISTEN_ON)
	log.Fatal(http.ListenAndServe(config.LISTEN_ON, r))
}
