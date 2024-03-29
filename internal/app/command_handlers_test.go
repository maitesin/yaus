package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestCreateShortenedURLHandler_Handle(t *testing.T) {
	repositorySaveError := errors.New("something went wrong saving the URL")
	repositoryFindError := errors.New("something went wrong finding the URL")
	validCmdOriginal := "https://oscarforner.com"

	tests := []struct {
		name                    string
		urlsRepositorySaveError error
		urlsRepositoryFindError error
		cmd                     app.Command
		expectedErr             error
	}{
		{
			name: `Given all valid inputs,
				   when the CreateShortenedURL Handle method is called,
                   then it returns no error`,
			cmd: app.CreateShortenedURLCmd{
				Original: validCmdOriginal,
			},
		},
		{
			name: `Given an invalid original URL input,
				   when the CreateShortenedURL Handle method is called,
                   then it returns an OriginalURLInvalidError`,
			cmd: app.CreateShortenedURLCmd{
				Original: "",
			},
			expectedErr: domain.NewOriginalURLInvalidError(""),
		},
		{
			name: `Given a URLsRepository that fails when the Save method is called,
				   when the CreateShortenedURL Handle method is called,
                   then it returns the error from the Save method`,
			urlsRepositorySaveError: repositorySaveError,
			cmd: app.CreateShortenedURLCmd{
				Original: validCmdOriginal,
			},
			expectedErr: repositorySaveError,
		},
		{
			name: `Given a URLsRepository that fails when the FindByOriginal method is called with an error not found,
				   when the CreateShortenedURL Handle method is called,
                   then it returns the error from the FindByOriginal method`,
			urlsRepositoryFindError: app.URLNotFound{},
			cmd: app.CreateShortenedURLCmd{
				Original: validCmdOriginal,
			},
			expectedErr: app.URLNotFound{},
		},
		{
			name: `Given a URLsRepository that fails when the FindByOriginal method is called with an error,
				   when the CreateShortenedURL Handle method is called,
                   then it returns the error from the FindByOriginal method`,
			urlsRepositoryFindError: repositoryFindError,
			cmd: app.CreateShortenedURLCmd{
				Original: validCmdOriginal,
			},
			expectedErr: repositoryFindError,
		},
		{
			name: `Given an invalid command,
				   when the CreateShortenedURL Handle method is called,
                   then it returns an InvalidCommandError`,
			cmd: &CommandMock{
				NameFunc: func() string {
					return "another command"
				},
			},
			expectedErr: app.InvalidCommandError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			stringGenerator := &StringGeneratorMock{
				GenerateFunc: func() string {
					return "abcdefg"
				},
			}

			urlsRepository := &URLsRepositoryMock{
				SaveFunc: func(context.Context, domain.URL) error {
					return tt.urlsRepositorySaveError
				},
				FindByOriginalFunc: func(context.Context, string) (domain.URL, error) {
					return domain.URL{}, tt.urlsRepositoryFindError
				},
			}

			csuh := app.NewCreateShortenedURLHandler(stringGenerator, urlsRepository)
			err := csuh.Handle(context.Background(), tt.cmd)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
		})
	}
}
