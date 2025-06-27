package types

import "github.com/golang-jwt/jwt/v5"

type IDTokenClaims struct {
	jwt.RegisteredClaims
}
