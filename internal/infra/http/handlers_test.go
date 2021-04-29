package http_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	"github.com/stretchr/testify/require"
)

func TestNewCreateShortenedHandler(t *testing.T) {
	tests := []struct {
		name                 string
		bodyReader           io.Reader
		cmdHandlerErr        error
		queryHandlerResponse app.QueryResponse
		queryHandlerErr      error
		expectedStatusCode   int
	}{
		{
			name: `Given a CreateShortenedHandler with a working command handler,
                   when an HTTP request is received,
                   then it returns an OK status code`,
			bodyReader:           strings.NewReader(""),
			cmdHandlerErr:        nil,
			queryHandlerResponse: domain.URL{Shortened: "1234567890"},
			expectedStatusCode:   http.StatusCreated,
		},
		{
			name: `Given a CreateShortenedHandler,
                   when an HTTP request with an empty bodyReader is received,
                   then it returns a bad request status code`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: `Given a CreateShortenedHandler with a non-working command handler,
                   when an HTTP request is received,
                   then it returns an internal server error status code`,
			bodyReader:         strings.NewReader(""),
			cmdHandlerErr:      errors.New("something went wrong in the Handler"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: `Given a CreateShortenedHandler with a non-working query handler,
                   when an HTTP request is received,
                   then it returns an internal server error status code`,
			bodyReader:         strings.NewReader(""),
			queryHandlerErr:    errors.New("something went wrong in the Handler"),
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: `Given a CreateShortenedHandler with a working query handler that does not return domain.URL,
                   when an HTTP request is received,
                   then it returns an internal server error status code`,
			bodyReader:           strings.NewReader(""),
			queryHandlerResponse: "unexpected",
			expectedStatusCode:   http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/", tt.bodyReader)
			require.NoError(t, err)

			res := httptest.NewRecorder()

			cmdHandler := &CommandHandlerMock{
				HandleFunc: func(context.Context, app.Command) error {
					return tt.cmdHandlerErr
				},
			}
			queryHandler := &QueryHandlerMock{
				HandleFunc: func(context.Context, app.Query) (app.QueryResponse, error) {
					return tt.queryHandlerResponse, tt.queryHandlerErr
				},
			}
			renderer := &RendererMock{RenderFunc: func(_ http.ResponseWriter, statusCode int, _ http.Header, _ interface{}) {
				require.Equal(t, tt.expectedStatusCode, statusCode)
			}}

			httpx.NewCreateShortenedHandler(cmdHandler, queryHandler, renderer)(res, req)
		})
	}
}

func TestNewRetrieveURLHandler(t *testing.T) {
	tests := []struct {
		name               string
		shortened          string
		queryHandlerRes    app.QueryResponse
		queryHandlerErr    error
		expectedStatusCode int
		expectedHeaders    http.Header
	}{
		{
			name: `Given a RetrieveURLByShortenedHandler with a working query handler,
                   when an HTTP request with a valid shortened code is received,
                   then it returns a 307 status code with the original URL in the Location header`,
			shortened:          "wololo",
			queryHandlerRes:    domain.URL{Original: "https://oscarforner.com"},
			queryHandlerErr:    nil,
			expectedStatusCode: http.StatusTemporaryRedirect,
			expectedHeaders: http.Header{
				"Location": []string{"https://oscarforner.com"},
			},
		},
		{
			name: `Given a RetrieveURLByShortenedHandler with a working query handler,
                   when an HTTP request with an empty shortened code is received,
                   then it returns a not found status code`,
			shortened:          "",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: `Given a RetrieveURLByShortenedHandler with a non-working query handler,
                   when an HTTP request with a valid shortened code is received,
                   then it returns a not found status code`,
			shortened:          "wololo",
			queryHandlerErr:    errors.New(""),
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: `Given a RetrieveURLByShortenedHandler with a unexpected query handler,
                   when an HTTP request with a valid shortened code is received,
                   then it returns an internal server error status code`,
			shortened:          "wololo",
			queryHandlerRes:    &QueryResponseMock{},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			queryHandler := &QueryHandlerMock{
				HandleFunc: func(context.Context, app.Query) (app.QueryResponse, error) {
					return tt.queryHandlerRes, tt.queryHandlerErr
				},
			}

			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("shortened", tt.shortened)
			ctx := context.WithValue(context.Background(), chi.RouteCtxKey, chiCtx)

			renderer := &RendererMock{RenderFunc: func(_ http.ResponseWriter, statusCode int, headers http.Header, _ interface{}) {
				require.Equal(t, tt.expectedStatusCode, statusCode)
				require.Equal(t, tt.expectedHeaders, headers)
			}}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/something", nil)
			require.NoError(t, err)

			res := httptest.NewRecorder()

			httpx.NewRetrieveURLHandler(queryHandler, renderer)(res, req)
		})
	}
}
