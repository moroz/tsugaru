package config_test

import (
	"crypto/ed25519"
	"oauth-provider/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadKeypair(t *testing.T) {
	privKey := `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIBi1ncaiitVN5aBTpDq/zQ37KyZDsqnWYLx7ip/QrPJU
-----END PRIVATE KEY-----`
	key, err := config.LoadKeypair([]byte(privKey))
	assert.NoError(t, err)
	assert.IsType(t, ed25519.PrivateKey{}, key)
}
