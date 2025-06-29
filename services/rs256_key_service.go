package services

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"oauth-provider/types"

	"github.com/golang-jwt/jwt/v5"
)

type rs256KeyService struct {
	privKey *rsa.PrivateKey
}

func RS256KeyServiceFromPEM(pem []byte) (KeyService, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}

	return &rs256KeyService{key}, nil
}

func (s *rs256KeyService) VerificationKey(*jwt.Token) (any, error) {
	return s.privKey.Public(), nil
}

func base64Uint(v *big.Int) string {
	return base64.RawURLEncoding.EncodeToString(v.Bytes())
}

func (s *rs256KeyService) calculateThumbprint(key *rsa.PublicKey) string {
	base := fmt.Sprintf(`{"e":"%s","kty":"RSA","n":"%s"}`, base64Uint(big.NewInt(int64(key.E))), base64Uint(key.N))
	hash := sha256.Sum256([]byte(base))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func (s *rs256KeyService) JWKS() types.JSONWebKeySet {
	key := s.privKey.Public().(*rsa.PublicKey)

	return types.JSONWebKeySet{
		Keys: []types.JSONWebKey{
			{
				KeyType:   types.JWKKeyType_RSA,
				Algorithm: types.JWKAlgorithm_RS256,
				E:         base64Uint(big.NewInt(int64(key.E))),
				N:         base64Uint(key.N),
				KeyID:     s.calculateThumbprint(key),
			},
		},
	}
}
