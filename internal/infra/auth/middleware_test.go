package auth_test

import (
	"encoding/base64"
	"fmt"
	"github.com/maitesin/yaus/internal/infra/auth"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type authHeaderBuilder func(auth.Config) string

func workingAuthHeaderBuilder(conf auth.Config) string {
	authBase64 := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", conf.User, conf.Pass)),
	)

	return fmt.Sprintf("Basic %s", authBase64)
}

func TestMiddleware(t *testing.T) {
	successHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	authConfig := auth.Config{
		User:    "wololo",
		Pass:    "42",
		Enabled: true,
	}

	tests := []struct {
		name               string
		authHeaderBuilder  authHeaderBuilder
		expectedStatusCode int
	}{
		{
			name: `Given a request containing a valid Authorization header,
                   when the authentication middleware is called,
				   then the provided handler is called`,
			authHeaderBuilder:  workingAuthHeaderBuilder,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: `Given a request containing an invalid Authorization header,
                   when the authentication middleware is called,
				   then an unauthorized status code is returned`,
			authHeaderBuilder: func(conf auth.Config) string {
				return workingAuthHeaderBuilder(conf) + "INVALID"
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			name: `Given a request containing a valid Authorization header, but with incorrect values,
                   when the authentication middleware is called,
				   then an unauthorized status code is returned`,
			authHeaderBuilder: func(conf auth.Config) string {
				conf.User = "something else"
				return workingAuthHeaderBuilder(conf)
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(http.MethodPost, "/", nil)
			require.NoError(t, err)
			req.Header.Add("Authorization", tt.authHeaderBuilder(authConfig))
			res := httptest.NewRecorder()

			middleware := auth.Middleware(authConfig)
			middleware(successHandler)(res, req)

			require.Equal(t, tt.expectedStatusCode, res.Code)
		})
	}
}
