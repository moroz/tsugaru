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

const privKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCoZQXdi4a/nlUw
oUtjK+sQd8L3XKL5khoWJZZU0Hfyc2I24ja+3pylQe+Q1xQJAy59tOWWl2zTCFgT
WlQyn6BpT9TfInA8w3RJrd0Y4YtDR7GusDtZWEpwtWpD1RXEMxHVaTweceYXm6pI
ENETOUBzygUuN40fTl87L6AxDaFzA7BoHb+9rKecU1cCt7iSk/z0Y5IkNo/fM1bb
giELLBbhQVLdQIqHNylWc5Fus6CJkdjXaGE87rHMIkK2Q2NuZsvy5DCejVi1hmbu
iGmQ+Il8QXhxWCpdAirFDS6RjyD3OGQyNX6d5TysgF9MZXRd+bK9Wqe96tBfBn6p
VC8HYmnVAgMBAAECggEACzPDPuB6C5PFH4nPWc7RYaRR7kI1nkwjm58O/9/pZtHZ
sgR891gYTB8ViIliThIt3NN0pX630NcggtMSwFZhpbfXlat1E3m+Keucxnu0l7p0
fcZAAHrQ4uwSoGYTv1xVXqDUTMMvdxkWLBqgSfrYSIujKhasdu1wBOAvdCvvtlan
4QLxFiokmIuHNNxLk20fqHWV9aJUsZk6RlRDN3ERMAKf+n3VGWXM/VMVHd4k0FCU
t7+4kel85za2yS63XL5Y0NMx7yb525iyS/yAqG3EFxjCtxDkgee/A1GPjHU1Eh7E
50QJdr5VCvLpyII+yTkpi0BFAt/iYtgy+V0xSIRtmQKBgQDselJBXxATklAgBxWE
0nO+0NVyx085YMPLOJ/XVvu3lFLT6EejnVTv/KiMlt2iVDVTihxMMvGxJCA5txry
fKoFAq6aDES+9oz4NDhhTimAetB/4//fFHiQqfBs81WzZ3O2zakCGRqiYDaveWvv
pVK3UUFOXu1frcWK9GGRxcc62QKBgQC2S9gGPS/jjFjA59hfxtZp0igHmSHchilx
RKNHnz0uGqoeTHW/dKR5FTR/DgwuERvPfRkYbBbwFfAJrZUgmrukri3hefbsAcsY
SGFYjGrcCPOO6YcALdH2laAHxX4GVEBWi511FsdbO2Tf7yh57vwq6pNwlgBJljkq
7aVRGaKxXQKBgQCTxWhLM1U/dameKe8Xfc8YSTVosQVfrIkNH8g6bz+CmywbAUZr
BnDCOpc0qz26J3bfSimesCL598Ivbq1nI+G2mdNzrgDzd+vlWfR1Ubt7bsOFd3s3
8nnYpGj6HCDMp/PWIrPe/ML4/riNdImvShbjOOJfT9Bzfoo357hkuDtkGQKBgQCa
AIa19pjvTdBo3zQu7WaTrUO974LorrpyAv9BcWgY+9O9lvBeVqbf16cqsu5dOHzb
E57Qv/e8yXuoYWk7Sxy8aZ0+/283P+iYUgVS7gUUb8d6cxRmdU8MVqkEB7aImEJm
GrphgWXXT9zPRVZXdCq6AsOd+Eqz3+HZvzvKwLJtzQKBgQDaN9E5LJgGtUqFsOrt
KKUNdWKhC+zRlN3dgv6NqED2chQs8JYpjhis+tpcfoDhwQAp4NEUYdnb8DZ4fp8+
gLhKdqX37rL3c84PJZlOSdQ533HHR4H+muId9AQfu3v9j6530cO4o2sl7wFEOrLn
EnFH8Z+kfvtwUESVnKeUeFtTFA==
-----END PRIVATE KEY-----`

func TestIDTokenClaims(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privKey))
	require.NoError(t, err)

	now := time.Now()
	claims := types.IDTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.BASE_URL,
			Subject:   "test",
			Audience:  []string{"test"},
			ExpiresAt: jwt.NewNumericDate(now.Add(60 * time.Minute)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	assert.NoError(t, err)
	assert.NotEqual(t, "", token)

	actual, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return key.Public(), nil
	})
	assert.NoError(t, err)
	iss, err := actual.Claims.GetIssuer()
	assert.NoError(t, err)
	assert.Equal(t, config.BASE_URL, iss)
}
