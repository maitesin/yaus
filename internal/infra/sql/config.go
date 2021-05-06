package sql

import (
	"fmt"
	"net"
	"net/url"
)

// Config defines the SQL connection configuration
type Config struct {
	Type         string
	Database     string
	User         string
	Password     string
	Host         string
	Port         string
	SSLMode      string
	BinaryParams string
}

// DatabaseURL returns the url prepared with its param values.
func (c *Config) DatabaseURL() string {
	u := url.URL{
		Scheme:   c.Type,
		User:     url.UserPassword(c.User, c.Password),
		Host:     net.JoinHostPort(c.Host, c.Port),
		Path:     c.Database,
		RawQuery: fmt.Sprintf("sslmode=%s&binary_parameters=%s", c.SSLMode, c.BinaryParams),
	}

	return u.String()
}
