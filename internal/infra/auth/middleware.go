package auth

import (
	"net/http"

	"github.com/maitesin/yaus/internal/infra/html"

	httpx "github.com/maitesin/yaus/internal/infra/http"
)

func Middleware(conf Config, renderer html.Renderer) httpx.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if conf.Enabled {
				user, pass, ok := request.BasicAuth()
				if !ok {
					renderer.Render(writer, http.StatusUnauthorized, nil, html.RendererValues{
						Shortened: "Unauthorized request",
						Category:  "danger",
					})
					return
				}

				if user != conf.User || pass != conf.Pass {
					renderer.Render(writer, http.StatusUnauthorized, nil, html.RendererValues{
						Shortened: "Unauthorized request",
						Category:  "danger",
					})
					return
				}
			}

			next(writer, request)
		}
	}
}
