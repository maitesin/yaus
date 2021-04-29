package http

import (
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
			buildResponse(w, http.StatusBadRequest, nil, []byte(InvalidRequestFormat))
			return
		}

		original := r.FormValue("url")
		cmd := app.CreateShortenedURLCmd{Original: original}
		err = commandHandler.Handle(r.Context(), cmd)
		if err != nil {
			renderer.Render(w, http.StatusInternalServerError, nil)
			return
		}

		query := app.RetrieveURLByOriginalQuery{Original: original}
		queryResponse, err := queryHandler.Handle(r.Context(), query)
		if err != nil {
			renderer.Render(w, http.StatusInternalServerError, nil)
			return
		}
		resp, ok := queryResponse.(domain.URL)
		if !ok {
			renderer.Render(w, http.StatusInternalServerError, nil)
			return
		}

		renderer.Render(w, http.StatusCreated, html.RendererValues{
			Shortened: resp.Shortened,
			Category:  "info",
		})
	}
}

func NewRetrieveURLHandler(handler app.QueryHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortened := chi.URLParam(r, "shortened")
		if shortened == "" {
			buildResponse(w, http.StatusNotFound, nil, []byte(NotFoundError))
			return
		}

		query := app.RetrieveURLByShortenedQuery{Shortened: shortened}
		response, err := handler.Handle(r.Context(), query)
		if err != nil {
			buildResponse(w, http.StatusNotFound, nil, []byte(NotFoundError))
			return
		}

		url, ok := response.(domain.URL)
		if !ok {
			buildResponse(w, http.StatusInternalServerError, nil, []byte(InternalServerError))
			return
		}

		header := http.Header{"Location": []string{url.Original}}
		buildResponse(w, http.StatusTemporaryRedirect, header, nil)
	}
}

func buildResponse(w http.ResponseWriter, status int, headers http.Header, body []byte) {
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(status)
	w.Write(body) //nolint:errcheck
}
