package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maitesin/yaus/config"
	"github.com/maitesin/yaus/internal/app"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	sqlx "github.com/maitesin/yaus/internal/infra/sql"
)

func main() {
	router := chi.NewRouter()
	conf := config.NewConfig()

	stringGenerator := app.NewRandomStringGenerator(&app.TimeProviderUTC{}, conf.RandomStringSize)
	urlsRepository := sqlx.NewInMemoryURLsRepository()

	router.Use(middleware.DefaultLogger)
	router.Post("/", httpx.NewCreateShortenedHandler(app.NewCreateShortenedURLHandler(stringGenerator, urlsRepository)))
	router.Get("/{shortened}", httpx.NewRetrieveURLHandler(app.NewRetrieveURLHandler(urlsRepository)))

	err := http.ListenAndServe(conf.HTTP.Address, router)
	if err != nil {
		fmt.Printf("Failed to start service: %s", err.Error())
	}
}
