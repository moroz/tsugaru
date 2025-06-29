package services

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"oauth-provider/types"

	"github.com/golang-jwt/jwt/v5"
)

type KeyService interface {
	// VerificationKey is a convenience function for use with jwt.Parse
	VerificationKey(*jwt.Token) (any, error)
	JWKS() types.JSONWebKeySet
}

type ed25519KeyService struct {
	privKey ed25519.PrivateKey
}

func Ed25519KeyService(privKey ed25519.PrivateKey) KeyService {
	return &ed25519KeyService{privKey}
}

func (s *ed25519KeyService) VerificationKey(*jwt.Token) (any, error) {
	return s.privKey.Public(), nil
}

func (s *ed25519KeyService) calculateThumbprint(key ed25519.PublicKey) string {
	base := fmt.Sprintf(`{"crv":"Ed25519","kty":"OKP","x":"%s"}`, base64.RawURLEncoding.EncodeToString(key))
	hash := sha256.Sum256([]byte(base))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func (s *ed25519KeyService) JWKS() types.JSONWebKeySet {
	key := s.privKey.Public().(ed25519.PublicKey)

	return types.JSONWebKeySet{Keys: []types.JSONWebKey{
		{KeyType: types.JWKKeyType_OKP,
			Curve:     types.JWKCurve_Ed25519,
			Use:       types.JWKUse_Signature,
			X:         base64.RawURLEncoding.EncodeToString(key),
			KeyID:     s.calculateThumbprint(key),
			Algorithm: types.JWKAlgorithm_EdDSA,
		}}}
}
