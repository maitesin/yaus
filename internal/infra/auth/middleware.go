package auth

import (
	httpx "github.com/maitesin/yaus/internal/infra/http"
	"net/http"
)

func Middleware(conf Config) httpx.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if conf.Enabled {
				user, pass, ok := request.BasicAuth()
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					return
				}

				if user != conf.User || pass != conf.Pass {
					writer.WriteHeader(http.StatusUnauthorized)
					return
				}
			}

			next(writer, request)
		}
	}
}
