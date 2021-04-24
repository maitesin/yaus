package config

import "github.com/maitesin/yaus/infra/http"

const (
	defaultAddress = ":8080"
	defaultSize    = 12
)

// Config defines the configuration of the YAUS application
type Config struct {
	HTTP             http.Config
	RandomStringSize int
}

// NewConfig is the constructor for the YAUS application configuration
func NewConfig() Config {
	return Config{
		HTTP: http.Config{
			Address: defaultAddress,
		},
		RandomStringSize: defaultSize,
	}
}
