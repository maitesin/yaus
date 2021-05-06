package config

import (
	"os"
	"path"
	"strconv"

	"github.com/maitesin/yaus/internal/infra/html"
	httpx "github.com/maitesin/yaus/internal/infra/http"
	"github.com/maitesin/yaus/internal/infra/sql"
)

const (
	defaultAddress    = ":8080"
	defaultAssetsPath = "devops/html"
	defaultSize       = "12"
)

// Config defines the configuration of the YAUS application
type Config struct {
	HTTP             httpx.Config
	HTML             html.Config
	SQL              sql.Config
	RandomStringSize int
}

// NewConfig is the constructor for the YAUS application configuration
func NewConfig() (Config, error) {
	size, err := strconv.Atoi(GetEnvOrDefault("YAUS_RANDOM_STRING_SIZE", defaultSize))
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTP: httpx.Config{
			Address: GetEnvOrDefault("YAUS_ADDRESS", defaultAddress),
		},
		HTML: html.Config{
			TemplatesDir: path.Join(GetEnvOrDefault("YAUS_ASSETS", defaultAssetsPath), "templates"),
			StaticDir:    path.Join(GetEnvOrDefault("YAUS_ASSETS", defaultAssetsPath), "static"),
		},
		SQL: sql.Config{
			Type:         GetEnvOrDefault("YAUS_DB_TYPE", "postgres"),
			Database:     GetEnvOrDefault("YAUS_DB_NAME", "urls"),
			User:         GetEnvOrDefault("YAUS_DB_USER", "yaus"),
			Password:     GetEnvOrDefault("YAUS_DB_PASSWORD", "postgres"),
			Host:         GetEnvOrDefault("YAUS_DB_HOST", "localhost"),
			Port:         GetEnvOrDefault("YAUS_DB_PORT", "54321"),
			SSLMode:      GetEnvOrDefault("YAUS_DB_SSL_MODE", "disable"),
			BinaryParams: GetEnvOrDefault("YAUS_DB_BINARY_PARAMETERS", "yes"),
		},
		RandomStringSize: size,
	}, nil
}

func GetEnvOrDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}

	return defaultValue
}
