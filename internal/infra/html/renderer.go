package html

import (
	"fmt"
	"html/template"
	"io"
	"path"
	"strings"
)

type Renderer struct {
	templatesDir string
}

func NewRenderer(templatesDir string) Renderer {
	return Renderer{templatesDir: templatesDir}
}

type RendererValues struct {
	Category  string
	Shortened string
}

func (hr Renderer) Render(writer io.Writer, names []string, values interface{}) {
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
}
