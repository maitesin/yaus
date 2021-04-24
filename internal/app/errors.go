package app

import "fmt"

const (
	errMsgInvalidCommand = "invalid command %q received. Expected %q"
	errMsgInvalidQuery   = "invalid query %q received. Expected %q"
	errMsgURLNotFound    = "url %q not found"
)

type InvalidCommandError struct {
	received Command
	expected Command
}

func (ice InvalidCommandError) Error() string {
	return fmt.Sprintf(errMsgInvalidCommand, ice.received.Name(), ice.expected.Name())
}

type InvalidQueryError struct {
	received Query
	expected Query
}

func (iqe InvalidQueryError) Error() string {
	return fmt.Sprintf(errMsgInvalidQuery, iqe.received.Name(), iqe.expected.Name())
}

type URLNotFound struct {
	url string
}

func NewURLNotFound(url string) URLNotFound {
	return URLNotFound{url: url}
}

func (unf URLNotFound) Error() string {
	return fmt.Sprintf(errMsgURLNotFound, unf.url)
}
