package config

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrInvalidKeypair = errors.New("invalid PEM key")

func LoadKeyPairFromFile(path string) (ed25519.PrivateKey, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadKeypair(bytes)
}

func Must[T any](v T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func LoadKeypair(pemBytes []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, ErrInvalidKeypair
	}

	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch k := parseResult.(type) {
	case ed25519.PrivateKey:
		return k, nil
	default:
		return nil, fmt.Errorf("%w (got type %T)", ErrInvalidKeypair, k)
	}
}
