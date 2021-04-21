package domain

import (
	"errors"
	"fmt"
)

const errOriginalURLInvalidMsg = "invalid URL %q"

// ErrShortenedValueIsEmpty is self explanatory
var ErrShortenedValueIsEmpty = errors.New("empty shortened value")

// OriginalURLInvalidError is self explanatory
type OriginalURLInvalidError struct {
	url string
}

// NewOriginalURLInvalidError is a constructor
func NewOriginalURLInvalidError(url string) OriginalURLInvalidError {
	return OriginalURLInvalidError{url: url}
}

// Error implements the error interface
func (ouie OriginalURLInvalidError) Error() string {
	return fmt.Sprintf(errOriginalURLInvalidMsg, ouie.url)
}
