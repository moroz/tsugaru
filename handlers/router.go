package handlers

import (
	"fmt"
	"net/http"
	"oauth-provider/db/queries"
	"oauth-provider/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func Router(db queries.DBTX, keyService services.KeyService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/", handleGet)

	discovery := DiscoveryController(keyService)
	r.Get("/.well-known/oauth/openid/jwks", discovery.Certs)
	r.Get("/.well-known/openid-configuration", discovery.OpenIDConfiguration)
	return r
}
