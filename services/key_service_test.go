package services_test

import (
	"oauth-provider/config"
	"oauth-provider/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEd25119KeyService(t *testing.T) {
	privKey := `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIBi1ncaiitVN5aBTpDq/zQ37KyZDsqnWYLx7ip/QrPJU
-----END PRIVATE KEY-----`
	key, err := config.LoadKeypair([]byte(privKey))
	require.NoError(t, err)

	ks := services.Ed25519KeyService(key)

	t.Run("generates JWKS", func(t *testing.T) {
		jwks := ks.JWKS()
		jwk := jwks.Keys[0]
		assert.Equal(t, "mnCbCFjL2LZAWAuX1qt2vp6AbNxT60_EMLQd0znNnPY", jwk.X)
		assert.Equal(t, "JTiwZ2nEHRp46Cw49NSbSWY4H88qo0fCJFfc74HuE94", jwk.KeyID)
		assert.Equal(t, "EdDSA", string(jwk.Algorithm))
		assert.Equal(t, "Ed25519", string(jwk.Curve))
		assert.Equal(t, "OKP", string(jwk.KeyType))
	})
}
