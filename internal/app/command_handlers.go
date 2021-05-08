package app

import (
	"context"
	"errors"

	"github.com/maitesin/yaus/internal/domain"
)

const (
	createShortenedURLCmdName = "CreateShortenedURLCmdName"
)

//go:generate moq -out zmock_command_test.go -pkg app_test . Command

// Command defines the interface of the commands to be performed
type Command interface {
	Name() string
}

//go:generate moq -out ../infra/http/zmock_command_test.go -pkg http_test . CommandHandler

// CommandHandler defines the interface of the handler to run commands
type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) error
}

// CreateShortenedURLCmd is a VTO
type CreateShortenedURLCmd struct {
	Original string
}

// Name returns the name of command to create a shortened URL
func (csum CreateShortenedURLCmd) Name() string {
	return createShortenedURLCmdName
}

// CreateShortenedURLHandler is the handler to create a shortened URL
type CreateShortenedURLHandler struct {
	stringGenerator StringGenerator
	urlsRepository  URLsRepository
}

// NewCreateShortenedURLHandler is a constructor
func NewCreateShortenedURLHandler(stringGenerator StringGenerator, urlsRepository URLsRepository) CreateShortenedURLHandler {
	return CreateShortenedURLHandler{stringGenerator: stringGenerator, urlsRepository: urlsRepository}
}

// Handle creates a shortened URL
func (csuh CreateShortenedURLHandler) Handle(ctx context.Context, cmd Command) error {
	createCmd, ok := cmd.(CreateShortenedURLCmd)
	if !ok {
		return InvalidCommandError{expected: CreateShortenedURLCmd{}, received: cmd}
	}

	_, err := csuh.urlsRepository.FindByOriginal(ctx, createCmd.Original)
	switch {
	case err == nil:
		return nil // The original URL already exists
	case !errors.As(err, &URLNotFound{}):
		return err
	}

	shortened := csuh.stringGenerator.Generate()
	url, err := domain.NewURL(createCmd.Original, shortened)
	if err != nil {
		return err
	}
	return csuh.urlsRepository.Save(ctx, url)
}
