package config

import (
	"log"
	"os"
)

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is not set!", key)
	}
	return val
}

var LISTEN_ON = ":3000"
var DATABASE_URL = MustGetenv("DATABASE_URL")
