package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/maitesin/yaus/internal/infra/html"
)

// NewCreateShortenedHandler returns an HTTP handler to process the creation of a shortened URL
func NewCreateShortenedHandler(
	commandHandler app.CommandHandler,
	queryHandler app.QueryHandler,
	renderer html.Renderer,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusBadRequest, nil, nil)
			return
		}

		original := r.FormValue("url")
		cmd := app.CreateShortenedURLCmd{Original: original}
		err = commandHandler.Handle(r.Context(), cmd)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusInternalServerError, nil, nil)
			return
		}

		query := app.RetrieveURLByOriginalQuery{Original: original}
		queryResponse, err := queryHandler.Handle(r.Context(), query)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusInternalServerError, nil, nil)
			return
		}
		resp, ok := queryResponse.(domain.URL)
		if !ok {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusInternalServerError, nil, nil)
			return
		}

		renderer.Render(w, http.StatusCreated, nil, html.RendererValues{
			Shortened: fmt.Sprintf("https://%s%s/%s", r.Host, r.RequestURI, resp.Shortened),
			Category:  "info",
		})
	}
}

func NewRetrieveURLHandler(handler app.QueryHandler, renderer html.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortened := chi.URLParam(r, "shortened")
		if shortened == "" {
			renderer.Render(w, http.StatusNotFound, nil, nil)
			return
		}

		query := app.RetrieveURLByShortenedQuery{Shortened: shortened}
		response, err := handler.Handle(r.Context(), query)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusNotFound, nil, nil)
			return
		}

		url, ok := response.(domain.URL)
		if !ok {
			fmt.Printf("Error: %s\n", err)
			renderer.Render(w, http.StatusInternalServerError, nil, nil)
			return
		}

		header := http.Header{"Location": []string{url.Original}}
		renderer.Render(w, http.StatusTemporaryRedirect, header, nil)
	}
}
