package main

import (
	"net/http"
	"platform/pkg/config"
	"platform/pkg/handlers"

	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	p := pat.New()
	p.Get("/:id", http.HandlerFunc(handlers.Repo.Get))
	p.Post("/:id", http.HandlerFunc(handlers.Repo.Set))
	return p
}
