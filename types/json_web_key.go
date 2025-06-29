package types

type JWKKeyType string

const (
	JWKKeyType_OKP JWKKeyType = "OKP"
	JWKKeyType_EC  JWKKeyType = "EC"
	JWKKeyType_RSA JWKKeyType = "RSA"
)

type JWKCurve string

const (
	JWKCurve_Ed25519 JWKCurve = "Ed25519"
	JWKCurve_Ed448   JWKCurve = "Ed448"
	JWKCurve_P256    JWKCurve = "P-256"
	JWKCurve_P384    JWKCurve = "P-384"
	JWKCurve_P521    JWKCurve = "P-521"
)

type JWKUse string

const (
	JWKUse_Signature  JWKUse = "sig"
	JWKUse_Encryption JWKUse = "enc"
)

type JWKAlgorithm string

const (
	JWKAlgorithm_EdDSA JWKAlgorithm = "EdDSA"
	JWKAlgorithm_RS256 JWKAlgorithm = "RS256"
)

type JSONWebKey struct {
	Algorithm JWKAlgorithm `json:"alg,omitempty"`
	Curve     JWKCurve     `json:"crv,omitempty"`
	KeyID     string       `json:"kid,omitempty"`
	KeyType   JWKKeyType   `json:"kty,omitempty"`
	Use       JWKUse       `json:"use,omitempty"`
	D         string       `json:"d,omitempty"`
	E         string       `json:"e,omitempty"`
	N         string       `json:"n,omitempty"`
	X         string       `json:"x,omitempty"`
}

type JSONWebKeySet struct {
	Keys []JSONWebKey `json:"keys"`
}
