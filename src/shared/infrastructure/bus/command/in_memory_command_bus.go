package command

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain/bus/command"
)

type InMemoryCommandBus struct {
	subscribers map[string]command.CommandHandler
}

func NewInMemoryCommandBus() InMemoryCommandBus {
	return InMemoryCommandBus{make(map[string]command.CommandHandler)}
}

func (imc InMemoryCommandBus) Dispatch(cmd command.Command) error {
	ch, ok := imc.subscribers[cmd.Name()]
	if !ok {
		return command.ErrHandlerNotFound
	}
	return ch.Handle(cmd)
}

func (imc *InMemoryCommandBus) RegisterHandler(cmd command.Command, handler command.CommandHandler) error {
	cmdName := cmd.Name()
	if _, ok := imc.subscribers[cmdName]; ok {
		return command.ErrHandlerAlreadyRegistered
	}
	imc.subscribers[cmdName] = handler
	return nil
}
