package main

import (
	"Bookins-and-reservations/WebPage/pkg/config"
	"Bookins-and-reservations/WebPage/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	//TODO : Fix image loading. Not working
	fileServer := http.FileServer(http.Dir("./WebPage/static/"))
	mux.Handle("/WebPage/static/*", http.StripPrefix("/WebPage/static", fileServer))

	return mux
}
