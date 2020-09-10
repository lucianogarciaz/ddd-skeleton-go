package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// ID is an UUID type identifier.
type ID uuid.UUID

// NewID is self-described.
func NewID() ID {
	return ID(uuid.New())
}

// ParseID is self-described.
func ParseID(s string) (ID, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID(uuid), nil
}

// MustParseID parses a string and if has any errors will panic.
func MustParseID(s string) ID {
	id, err := ParseID(s)
	if err != nil {
		panic(fmt.Sprintf("ID %s parse: %s", s, err.Error()))
	}
	return id
}

// String is self-described.
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// IsEmpty is self-described
func (id ID) IsEmpty() bool {
	return id == ID{}
}
