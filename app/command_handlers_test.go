package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/maitesin/yaus/internal/domain"

	"github.com/maitesin/yaus/app"

	"github.com/stretchr/testify/require"
)

func TestCreateShortenedURLHandler_Handle(t *testing.T) {
	repositoryError := errors.New("something went wrong saving the URL")
	validCmdOriginal := "https://oscarforner.com"

	tests := []struct {
		name                    string
		urlsRepositorySaveError error
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
			urlsRepositorySaveError: repositoryError,
			cmd: app.CreateShortenedURLCmd{
				Original: validCmdOriginal,
			},
			expectedErr: repositoryError,
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
			}

			csuh := app.NewCreateShortenedURLHandler(stringGenerator, urlsRepository)
			err := csuh.Handle(context.Background(), tt.cmd)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
		})
	}
}
