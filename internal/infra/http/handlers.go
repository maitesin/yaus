package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
)

// NewCreateShortenedHandler returns an HTTP handler to process the creation of a shortened URL
func NewCreateShortenedHandler(commandHandler app.CommandHandler, queryHandler app.QueryHandler) http.HandlerFunc {
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
			buildResponse(w, http.StatusInternalServerError, nil, []byte(InternalServerError))
			return
		}

		query := app.RetrieveURLByOriginalQuery{Original: original}
		queryResponse, err := queryHandler.Handle(r.Context(), query)
		if err != nil {
			buildResponse(w, http.StatusInternalServerError, nil, []byte(InternalServerError))
			return
		}
		resp, ok := queryResponse.(domain.URL)
		if !ok {
			buildResponse(w, http.StatusInternalServerError, nil, []byte(InternalServerError))
			return
		}

		buildResponse(w, http.StatusOK, nil, []byte(fmt.Sprintf("shortcode: %q", resp.Shortened)))
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
