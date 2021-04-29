package html

import (
	"fmt"
	"net/http"
)

//go:generate moq -out ../http/zmock_renderer_test.go -pkg http_test . Renderer

type Renderer interface {
	Render(writer http.ResponseWriter, httpStatus int, headers http.Header, values interface{})
}

type BasicRenderer struct {
	factory TemplateFactory
}

func NewBasicRenderer(factory TemplateFactory) BasicRenderer {
	return BasicRenderer{factory: factory}
}

type RendererValues struct {
	Category  string
	Shortened string
}

func (hr BasicRenderer) Render(writer http.ResponseWriter, httpStatus int, headers http.Header, values interface{}) {
	// response HTTP headers
	for key, values := range headers {
		for _, value := range values {
			writer.Header().Add(key, value)
		}
	}
	writer.WriteHeader(httpStatus)

	// response HTTP status code
	writer.WriteHeader(httpStatus)

	// response HTTP body
	template := hr.factory.Template(httpStatus)
	err := template.Execute(writer, values)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
}
