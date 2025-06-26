package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

var LISTEN_ON = ":3000"
var DATABASE_URL = MustGetenv("DATABASE_URL")

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is not set!", key)
	}
	return val
}

func main() {
	config, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	tracer := tracelog.TraceLog{}
	tracer.Logger = tracelog.LoggerFunc(logQuery)
	tracer.LogLevel = tracelog.LogLevelInfo
	config.ConnConfig.Tracer = &tracer

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	conn.Exec(context.Background(), "select now()")

	r := chi.NewRouter()
	r.Get("/", handleGet)
	log.Printf("Listening on %s", LISTEN_ON)
	log.Fatal(http.ListenAndServe(LISTEN_ON, r))
}
