package config

import (
	"github.com/maitesin/yaus/internal/infra/html"
	httpx "github.com/maitesin/yaus/internal/infra/http"
)

const (
	defaultAddress      = ":8080"
	defaultTemplatesDir = "devops/html/templates/"
	defaultStaticDir    = "./devops/html/static/"
	defaultSize         = 12
)

// Config defines the configuration of the YAUS application
type Config struct {
	HTTP             httpx.Config
	HTML             html.Config
	RandomStringSize int
}

// NewConfig is the constructor for the YAUS application configuration
func NewConfig() Config {
	return Config{
		HTTP: httpx.Config{
			Address: defaultAddress,
		},
		HTML: html.Config{
			TemplatesDir: defaultTemplatesDir,
			StaticDir:    defaultStaticDir,
		},
		RandomStringSize: defaultSize,
	}
}
