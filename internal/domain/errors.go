package domain

import "fmt"

const (
	errOriginalURLInvalidMsg    = "invalid URL %q"
	errShortenedValueIsEmptyMsg = "empty shortened value"
)

type OriginalURLInvalidError struct {
	url string
}

func NewOriginalURLInvalidError(url string) OriginalURLInvalidError {
	return OriginalURLInvalidError{url: url}
}

func (ouie OriginalURLInvalidError) Error() string {
	return fmt.Sprintf(errOriginalURLInvalidMsg, ouie.url)
}

type ShortenedValueIsEmptyError struct{}

func (sviee ShortenedValueIsEmptyError) Error() string {
	return errShortenedValueIsEmptyMsg
}
