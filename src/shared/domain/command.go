package domain

import "errors"

var ErrInvalidCommand = errors.New("invalid command")

// Command is the interface for identifying commands by name.
type Command interface {
	Name() string
}
