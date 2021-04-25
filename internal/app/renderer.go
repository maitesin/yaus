package app

import "io"

type Renderer interface {
	Render(writer io.Writer, names []string, values map[string]interface{})
}
