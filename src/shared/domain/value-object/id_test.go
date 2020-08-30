package value_object

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseID(t *testing.T) {
	require := require.New(t)
	id := NewID()
	parsedID, err := ParseID(id.String())
	require.NoError(err)
	require.IsType(ID{}, parsedID)
}

func TestMustParseID(t *testing.T) {
	require := require.New(t)
	t.Run(`Given a string id that isn't an UUID,
		when MustParseID is called,
		then it panics`, func(t *testing.T) {
		id := "should panic"
		require.Panics(func() { MustParseID(id) })
	})
	t.Run(`Given a string that is an UUID,
		when MustParseID is called,
		then it returns not panic and an ID must be returned`, func(t *testing.T) {
		stringID := NewID().String()
		require.NotPanics(func() {
			id := MustParseID(stringID)
			require.IsType(ID{}, id)
		})
	})
}

func TestIsEmpty(t *testing.T) {
	require := require.New(t)
	t.Run(`Given a non empty id,
		when IsEmpty is called,
		then it returns false`, func(t *testing.T) {
		id := NewID()
		require.False(id.IsEmpty())
	})
	t.Run(`Given an empty id,
		when IsEmpty is called,
		then it returns true`, func(t *testing.T) {
		id := ID{}
		require.True(id.IsEmpty())
	})
}
