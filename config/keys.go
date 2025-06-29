package config

import (
	"crypto/rsa"
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidKeypair = errors.New("invalid PEM key")

func LoadKeyPairFromFile(path string) (*rsa.PrivateKey, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(bytes)
}

func Must[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}
