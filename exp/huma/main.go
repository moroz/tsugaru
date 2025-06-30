package main

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type GreetingInput struct {
}

func main() {
	router := chi.NewRouter()
	api := humachi.New(router, huma.DefaultConfig("Authorizz API", "1.0.0"))

	huma.Get(api, "/greeting/{name}")
}
