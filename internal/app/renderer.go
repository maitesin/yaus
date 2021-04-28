package app

import "io"

//go:generate moq -out ../infra/http/zmock_renderer_test.go -pkg http_test . Renderer

type Renderer interface {
	Render(writer io.Writer, names []string, values interface{})
}
