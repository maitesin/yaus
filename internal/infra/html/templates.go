package html

import "net/http"

type TemplateFactory interface {
	Template(httpStatus int) []string
}

type YausTemplateFactory struct {
	status2templates map[int][]string
	defaultTemplates []string
}

func (ytf YausTemplateFactory) Template(httpStatus int) []string {
	found, ok := ytf.status2templates[httpStatus]
	if !ok {
		return ytf.defaultTemplates
	}
	return found
}

func NewYausTemplateFactory() YausTemplateFactory {
	return YausTemplateFactory{
		defaultTemplates: []string{"layout.html", "home.html"},
		status2templates: map[int][]string{
			http.StatusNotFound:            {"layout.html", "404.html"},
			http.StatusInternalServerError: {"layout.html", "500.html"},
		},
	}
}
