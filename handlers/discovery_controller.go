package handlers

import (
	"encoding/json"
	"net/http"
	"oauth-provider/config"
	"oauth-provider/services"
	"oauth-provider/types"
)

type discoveryController struct {
	ks services.KeyService
}

func DiscoveryController(ks services.KeyService) *discoveryController {
	return &discoveryController{ks}
}

func (c *discoveryController) Certs(w http.ResponseWriter, r *http.Request) {
	jwks := c.ks.JWKS()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jwks)
}

func (c *discoveryController) OpenIDConfiguration(w http.ResponseWriter, r *http.Request) {
	response := &types.OpenIDProviderMetadata{
		Issuer:                config.BASE_URL,
		AuthorizationEndpoint: config.BASE_URL + config.AUTHORIZATION_ENDPOINT,
		TokenEndpoint:         config.BASE_URL + config.TOKEN_ENDPOINT,
		JwksUri:               config.BASE_URL + config.JWKS_URI,
		SubjectTypesSupported: []string{
			"pairwise",
		},
		ScopesSupported: []string{
			"openid",
		},
		ResponseTypesSupported: []string{
			"code", "id_token", "code id_token", "id_token token",
		},
		IDTokenSigningAlgValuesSupported: []string{
			"RS256",
		},
		ClaimsSupported: []string{
			"sub", "iss", "aud", "exp", "iat", "email",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
