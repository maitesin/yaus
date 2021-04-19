package app

import "fmt"

const errMsgInvalidCommand = "invalid command %q received. Expected %q"
const errMsgURLNotFound = "url %q not found"

type InvalidCommandError struct {
	received Command
	expected Command
}

func (ice InvalidCommandError) Error() string {
	return fmt.Sprintf(errMsgInvalidCommand, ice.received.Name(), ice.expected.Name())
}

type URLNotFound struct {
	url string
}

func (unf URLNotFound) Error() string {
	return fmt.Sprintf(errMsgURLNotFound, unf.url)
}
