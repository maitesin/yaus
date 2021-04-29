package html

import (
	"html/template"
	"net/http"
	"path"
)

type TemplateFactory interface {
	Template(httpStatus int) *template.Template
}

type YausTemplateFactory struct {
	status2templates map[int]*template.Template
	defaultTemplates *template.Template
}

func (ytf YausTemplateFactory) Template(httpStatus int) *template.Template {
	found, ok := ytf.status2templates[httpStatus]
	if !ok {
		return ytf.defaultTemplates
	}
	return found
}

func NewYausTemplateFactory(templatesDir string) (*YausTemplateFactory, error) {
	defaultTemplate, err := buildTemplate(templatesDir, []string{"layout.html", "home.html"})
	if err != nil {
		return nil, err
	}
	statusNotFound, err := buildTemplate(templatesDir, []string{"layout.html", "404.html"})
	if err != nil {
		return nil, err
	}
	statusInternalServerError, err := buildTemplate(templatesDir, []string{"layout.html", "500.html"})
	if err != nil {
		return nil, err
	}
	return &YausTemplateFactory{
		defaultTemplates: defaultTemplate,
		status2templates: map[int]*template.Template{
			http.StatusNotFound:            statusNotFound,
			http.StatusInternalServerError: statusInternalServerError,
		},
	}, nil
}

func buildTemplate(templatesDir string, templateNames []string) (*template.Template, error) {
	templateFiles := make([]string, len(templateNames))
	for i, name := range templateNames {
		templateFiles[i] = path.Join(templatesDir, name)
	}

	parsed, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
