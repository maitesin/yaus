package http_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/maitesin/yaus/app"
	httpx "github.com/maitesin/yaus/infra/http"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestNewCreateShortenedHandler(t *testing.T) {
	tests := []struct {
		name               string
		body               string
		cmdHandlerErr      error
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name:               "",
			body:               "",
			cmdHandlerErr:      nil,
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte{},
		},
		{
			name:               "",
			body:               "",
			cmdHandlerErr:      errors.New("something went wrong in the Handler"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       []byte(httpx.InternalServerError),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/", strings.NewReader(tt.body))
			require.NoError(t, err)

			res := httptest.NewRecorder()

			cmdHandler := &CommandHandlerMock{
				HandleFunc: func(_ context.Context, _ app.Command) error {
					return tt.cmdHandlerErr
				},
			}
			httpx.NewCreateShortenedHandler(cmdHandler)(res, req)

			require.Equal(t, tt.expectedStatusCode, res.Code)
			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			require.Equal(t, tt.expectedBody, body)
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
		expectedBody       []byte
	}{
		{
			name:               "",
			shortened:          "wololo",
			queryHandlerRes:    domain.URL{Original: "https://oscarforner.com"},
			queryHandlerErr:    nil,
			expectedStatusCode: http.StatusTemporaryRedirect,
			expectedHeaders: http.Header{
				"Location": []string{"https://oscarforner.com"},
			},
			expectedBody: []byte{},
		},
		{
			name:               "",
			shortened:          "",
			expectedStatusCode: http.StatusNotFound,
			expectedHeaders:    http.Header{},
			expectedBody:       []byte(httpx.NotFoundError),
		},
		{
			name:               "",
			shortened:          "wololo",
			queryHandlerErr:    errors.New(""),
			expectedStatusCode: http.StatusNotFound,
			expectedHeaders:    http.Header{},
			expectedBody:       []byte(httpx.NotFoundError),
		},
		{
			name:               "",
			shortened:          "wololo",
			queryHandlerRes:    &QueryResponseMock{},
			expectedStatusCode: http.StatusInternalServerError,
			expectedHeaders:    http.Header{},
			expectedBody:       []byte(httpx.InternalServerError),
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

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/something", nil)
			require.NoError(t, err)

			res := httptest.NewRecorder()

			httpx.NewRetrieveURLHandler(queryHandler)(res, req)

			require.Equal(t, tt.expectedStatusCode, res.Code)
			require.Equal(t, tt.expectedHeaders, res.Header())
			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			require.Equal(t, tt.expectedBody, body)
		})
	}
}
