package handlers

import (
	"encoding/json"
	"net/http"
	"oauth-provider/services"
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

func OpenIDConfiguration(w http.ResponseWriter, r *http.Request) {

}
