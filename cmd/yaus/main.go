package main

import (
	"fmt"
	"net/http"

	"github.com/maitesin/yaus/config"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/infra/html"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	sqlx "github.com/maitesin/yaus/internal/infra/sql"
)

func main() {
	conf := config.NewConfig()

	stringGenerator := app.NewRandomStringGenerator(&app.TimeProviderUTC{}, conf.RandomStringSize)
	urlsRepository := sqlx.NewInMemoryURLsRepository()
	templateFactory, err := html.NewYausTemplateFactory(conf.HTML.TemplatesDir)
	if err != nil {
		fmt.Printf("Failed to create template factory: %s\n", err.Error())
		return
	}
	renderer := html.NewBasicRenderer(templateFactory)

	err = http.ListenAndServe(conf.HTTP.Address, httpx.DefaultRouter(conf.HTML, urlsRepository, stringGenerator, renderer))
	if err != nil {
		fmt.Printf("Failed to start service: %s\n", err.Error())
	}
}
