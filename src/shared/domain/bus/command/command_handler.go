package command

// CommandHandler is the interface for running logic based on a command input
type CommandHandler interface {
	Handle(Command) error
}
