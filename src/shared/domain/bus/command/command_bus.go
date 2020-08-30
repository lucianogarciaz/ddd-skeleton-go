package command

import "errors"

var (
	ErrHandlerAlreadyRegistered = errors.New("the command is already registered")
	ErrHandlerNotFound          = errors.New("the handler was not found")
)

// CommandBus is the interface for command bus
type CommandBus interface {
	Dispatch(Command) error
	RegisterHandler(Command, CommandHandler) error
}
