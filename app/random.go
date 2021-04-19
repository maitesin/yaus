package app

import "math/rand"

//go:generate moq -out mock_string_generator_test.go -pkg app_test . StringGenerator

// StringGenerator defines the interface to generate strings of a given length
type StringGenerator interface {
	Generate(size int) string
}

// RandomStringGenerator implements the StringGenerator interface to return random strings
type RandomStringGenerator struct {
	timeProvider TimeProvider
}

// NewRandomStringGenerator is a constructor
func NewRandomStringGenerator(timeProvider TimeProvider) RandomStringGenerator {
	return RandomStringGenerator{
		timeProvider: timeProvider,
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Generate returns a string of size characters randomly generated
func (rsg RandomStringGenerator) Generate(size int) string {
	rand.Seed(rsg.timeProvider.Now().UnixNano())
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))] //nolint: gosec
	}
	return string(b)
}
