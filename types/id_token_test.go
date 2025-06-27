package types_test

import (
	"oauth-provider/config"
	"oauth-provider/types"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIDTokenClaims(t *testing.T) {
	privKey := `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIBi1ncaiitVN5aBTpDq/zQ37KyZDsqnWYLx7ip/QrPJU
-----END PRIVATE KEY-----`
	key, err := config.LoadKeypair([]byte(privKey))
	require.NoError(t, err)

	now := time.Now()
	claims := types.IDTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "test",
			Subject:   "test",
			Audience:  []string{"test"},
			ExpiresAt: jwt.NewNumericDate(now.Add(60 * time.Minute)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token, err := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims).SignedString(key)
	assert.NoError(t, err)
	assert.NotEqual(t, "", token)

	actual, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return key.Public(), nil
	})
	assert.NoError(t, err)
	iss, err := actual.Claims.GetIssuer()
	assert.NoError(t, err)
	assert.Equal(t, "test", iss)
}
