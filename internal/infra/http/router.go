package http

import (
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/infra/html"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func DefaultRouter(conf html.Config, authMiddleware Middleware, repository app.URLsRepository, generator app.StringGenerator, renderer html.Renderer) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		renderer.Render(writer, http.StatusOK, nil, nil)
	})

	router.Post("/u", authMiddleware(NewCreateShortenedHandler(
		app.NewCreateShortenedURLHandler(generator, repository),
		app.NewRetrieveURLByOriginalHandler(repository),
		renderer,
	)))

	router.Get("/u/{shortened}", NewRetrieveURLHandler(app.NewRetrieveURLByShortenedHandler(repository), renderer))
	router.Get("/css/main.css", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, path.Join(conf.StaticDir, "main.css"))
	})
	router.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		renderer.Render(writer, http.StatusNotFound, nil, nil)
	})

	return router
}
