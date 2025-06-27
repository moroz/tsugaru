package main

import (
	"context"
	"log"
	"net/http"
	"oauth-provider/config"
	"oauth-provider/handlers"
	"oauth-provider/services"

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

	pk := config.Must(config.LoadKeyPairFromFile("keys/private.pem"))

	r := handlers.Router(conn, services.Ed25519KeyService(pk))

	log.Printf("Listening on %s", config.LISTEN_ON)
	log.Fatal(http.ListenAndServe(config.LISTEN_ON, r))
}
