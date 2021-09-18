package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/maitesin/yaus/config"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/infra/auth"
	"github.com/maitesin/yaus/internal/infra/html"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	sqlx "github.com/maitesin/yaus/internal/infra/sql"
	"github.com/upper/db/v4/adapter/postgresql"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to create the configuration: %s\n", err.Error())
		return
	}

	dbConn, err := sql.Open("postgres", conf.SQL.DatabaseURL())
	if err != nil {
		fmt.Printf("Failed to open connection to the DB: %s\n", err)
		return
	}
	defer dbConn.Close()

	pgConn, err := postgresql.New(dbConn)
	if err != nil {
		fmt.Printf("Failed to initialize connection with the DB: %s\n", err)
		return
	}
	defer pgConn.Close()

	urlsRepository := sqlx.NewURLsRepository(pgConn)
	stringGenerator := app.NewRandomStringGenerator(&app.TimeProviderUTC{}, conf.RandomStringSize)
	templateFactory, err := html.NewYausTemplateFactory(conf.HTML.TemplatesDir)
	if err != nil {
		fmt.Printf("Failed to create template factory: %s\n", err.Error())
		return
	}
	renderer := html.NewBasicRenderer(templateFactory)

	authMiddleware := auth.Middleware(conf.Auth)

	err = http.ListenAndServe(
		strings.Join([]string{conf.HTTP.Host, conf.HTTP.Port}, ":"),
		httpx.DefaultRouter(conf.HTML, authMiddleware, urlsRepository, stringGenerator, renderer),
	)
	if err != nil {
		fmt.Printf("Failed to start service: %s\n", err.Error())
	}
}
