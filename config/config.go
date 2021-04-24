package config

import (
	http2 "github.com/maitesin/yaus/internal/infra/http"
)

const (
	defaultAddress = ":8080"
	defaultSize    = 12
)

// Config defines the configuration of the YAUS application
type Config struct {
	HTTP             http2.Config
	RandomStringSize int
}

// NewConfig is the constructor for the YAUS application configuration
func NewConfig() Config {
	return Config{
		HTTP: http2.Config{
			Address: defaultAddress,
		},
		RandomStringSize: defaultSize,
	}
}
