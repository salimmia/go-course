package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/salimmia/go-course/pkg/config"
	"github.com/salimmia/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{
	mux := pat.New()
	
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	
	return mux
}