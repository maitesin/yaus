package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maitesin/yaus/config"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/infra/html"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	sqlx "github.com/maitesin/yaus/internal/infra/sql"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)

	conf := config.NewConfig()

	stringGenerator := app.NewRandomStringGenerator(&app.TimeProviderUTC{}, conf.RandomStringSize)
	urlsRepository := sqlx.NewInMemoryURLsRepository()
	templateFactory, err := html.NewYausTemplateFactory(conf.HTML.TemplatesDir)
	if err != nil {
		fmt.Printf("Failed to create template factory: %s\n", err.Error())
		return
	}
	renderer := html.NewBasicRenderer(templateFactory)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		renderer.Render(writer, http.StatusOK, nil)
	})
	router.Post("/u", httpx.NewCreateShortenedHandler(
		app.NewCreateShortenedURLHandler(stringGenerator, urlsRepository),
		app.NewRetrieveURLByOriginalHandler(urlsRepository),
		renderer,
	))
	router.Get("/u/{shortened}", httpx.NewRetrieveURLHandler(app.NewRetrieveURLByShortenedHandler(urlsRepository)))
	router.Get("/css/main.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join(conf.HTML.StaticDir, "main.css"))
	})
	router.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		renderer.Render(writer, http.StatusNotFound, nil)
	})

	err = http.ListenAndServe(conf.HTTP.Address, router)
	if err != nil {
		fmt.Printf("Failed to start service: %s", err.Error())
	}
}
