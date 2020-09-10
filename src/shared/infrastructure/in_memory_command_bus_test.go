package infrastructure_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/infrastructure"

	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

func TestCommandBus(t *testing.T) {
	cmd := &domain.CommandMock{
		NameFunc: func() string {
			return "some name"
		},
	}
	t.Run(`Given an in memory command bus and a no registered command
	when it's called,
	then it returns ErrHandlerNotFound`, func(t *testing.T) {
		cmdBus := infrastructure.NewInMemoryCommandBus()

		err := cmdBus.Dispatch(cmd)
		require.True(t, errors.As(err, &domain.ErrHandlerNotFound))
	})

	t.Run(`Given an in memory command bus and a registered command
	when it's called,
	then it the command handler is executed`, func(t *testing.T) {
		cmdHandler := &domain.CommandHandlerMock{
			HandleFunc: func(domain.Command) error {
				return nil
			},
		}
		cmdBus := infrastructure.NewInMemoryCommandBus()
		err := cmdBus.RegisterHandler(cmd, cmdHandler)
		require.NoError(t, err)
		err = cmdBus.Dispatch(cmd)
		require.NoError(t, err)
		require.Len(t, cmdHandler.HandleCalls(), 1)
	})
}
