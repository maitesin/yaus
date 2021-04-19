package app

import (
	"time"
)

//go:generate moq -out mock_time_provider_test.go -pkg app_test . TimeProvider

// TimeProvider defines the way to obtain the current time
type TimeProvider interface {
	Now() time.Time
}

// TimeProviderUTC provides time in UTC
type TimeProviderUTC struct{}

// Now returns the current time
func (tpu *TimeProviderUTC) Now() time.Time {
	return time.Now().UTC()
}
