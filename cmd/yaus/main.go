package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maitesin/yaus/app"
	"github.com/maitesin/yaus/config"
	httpx "github.com/maitesin/yaus/infra/http"
	"github.com/maitesin/yaus/infra/sql"
)

func main() {
	router := chi.NewRouter()
	conf := config.NewConfig()

	stringGenerator := app.NewRandomStringGenerator(&app.TimeProviderUTC{}, conf.RandomStringSize)
	urlsRepository := sql.NewInMemoryURLsRepository()

	router.Use(middleware.DefaultLogger)
	router.Post("/", httpx.NewCreateShortenedHandler(app.NewCreateShortenedURLHandler(stringGenerator, urlsRepository)))
	router.Get("/{shortened}", httpx.NewRetrieveURLHandler(app.NewRetrieveURLHandler(urlsRepository)))

	err := http.ListenAndServe(conf.HTTP.Address, router)
	if err != nil {
		fmt.Printf("Failed to start service: %s", err.Error())
	}
}
