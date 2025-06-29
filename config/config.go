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

func GetenvWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return def
}

var LISTEN_ON = ":3000"
var DATABASE_URL = MustGetenv("DATABASE_URL")
var HOST = GetenvWithDefault("HOST", "auth.authorizz.localhost")

var BASE_URL = "https://" + HOST

const AUTHORIZATION_ENDPOINT = "/oauth/authorize"
const TOKEN_ENDPOINT = "/oauth/token"
const JWKS_URI = "/.well-known/oauth/openid/jwks"
