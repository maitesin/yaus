package html

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
)

//go:generate moq -out ../http/zmock_renderer_test.go -pkg http_test . Renderer

type Renderer interface {
	Render(writer http.ResponseWriter, httpStatus int, values interface{})
}

type BasicRenderer struct {
	templatesDir string
	factory      TemplateFactory
}

func NewBasicRenderer(templatesDir string, factory TemplateFactory) BasicRenderer {
	return BasicRenderer{
		templatesDir: templatesDir,
		factory:      factory,
	}
}

type RendererValues struct {
	Category  string
	Shortened string
}

func (hr BasicRenderer) Render(writer http.ResponseWriter, httpStatus int, values interface{}) {
	names := hr.factory.Template(httpStatus)
	templateFiles := make([]string, len(names))
	for i, name := range names {
		templateFiles[i] = path.Join(hr.templatesDir, name)
	}

	parsed, err := template.ParseFiles(templateFiles...)
	if err != nil {
		fmt.Printf("Error parsing templates %q: %s\n", strings.Join(templateFiles, ","), err.Error())
		return
	}
	err = parsed.Execute(writer, values)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
	writer.WriteHeader(httpStatus)
}
