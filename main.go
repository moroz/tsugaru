package main

import (
	"context"
	"log"
	"net/http"
	"oauth-provider/config"
	"oauth-provider/handlers"
	"oauth-provider/services"
	"os"

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

func main() {
	cfg := config.Must(pgxpool.ParseConfig(config.DATABASE_URL))

	tracer := tracelog.TraceLog{}
	tracer.Logger = tracelog.LoggerFunc(logQuery)
	tracer.LogLevel = tracelog.LogLevelInfo
	cfg.ConnConfig.Tracer = &tracer

	db := config.Must(pgxpool.NewWithConfig(context.Background(), cfg))
	pem := config.Must(os.ReadFile("keys/private.pem"))
	r := handlers.Router(db, config.Must(services.RS256KeyServiceFromPEM(pem)))

	log.Printf("Listening on %s", config.LISTEN_ON)
	log.Fatal(http.ListenAndServe(config.LISTEN_ON, r))
}
