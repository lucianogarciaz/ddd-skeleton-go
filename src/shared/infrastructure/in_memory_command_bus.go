package infrastructure

import (
	"github.com/lucianogarciaz/ddd-skeleton-go/src/shared/domain"
)

type InMemoryCommandBus struct {
	subscribers map[string]domain.CommandHandler
}

func NewInMemoryCommandBus() InMemoryCommandBus {
	return InMemoryCommandBus{make(map[string]domain.CommandHandler)}
}

func (imc InMemoryCommandBus) Dispatch(cmd domain.Command) error {
	ch, ok := imc.subscribers[cmd.Name()]
	if !ok {
		return domain.ErrHandlerNotFound
	}
	return ch.Handle(cmd)
}

func (imc *InMemoryCommandBus) RegisterHandler(cmd domain.Command, handler domain.CommandHandler) error {
	cmdName := cmd.Name()
	if _, ok := imc.subscribers[cmdName]; ok {
		return domain.ErrHandlerAlreadyRegistered
	}
	imc.subscribers[cmdName] = handler
	return nil
}
