package app

import (
	"context"

	"github.com/maitesin/yaus/internal/domain"
)

const (
	createShortenedURLCmdName = "CreateShortenedURLCmdName"
)

// Command defines the interface of the commands to be performed
type Command interface {
	Name() string
}

// CommandHandler defines the interface of the handler to run commands
type CommandHandler interface {
	Handle(ctx context.Context, cmd Command) error
}

// CreateShortenedURLCmd is a VTO
type CreateShortenedURLCmd struct {
	original string
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

	shortened := csuh.stringGenerator.Generate()
	return csuh.urlsRepository.Save(ctx, domain.NewURL(createCmd.original, shortened))
}
