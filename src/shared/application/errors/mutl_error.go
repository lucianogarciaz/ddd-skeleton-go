package errors

import (
	"errors"
	"strings"
)

var _ error = &MultiError{}

const (
	multiErrorPrefix    = "multi error: "
	multiErrorSeparator = "; "
)

// NewMultiError returns a new MultiError
func NewMultiError() *MultiError {
	return &MultiError{
		errors: make([]error, 0),
	}
}

// MultiError is self described
type MultiError struct {
	errors []error
}

// Error complaints the Error interface
func (me MultiError) Error() string {
	if len(me.errors) == 0 {
		return ""
	}

	errorMsgs := make([]string, len(me.errors))
	for i, e := range me.errors {
		errorMsgs[i] = e.Error()
	}
	return multiErrorPrefix + strings.Join(errorMsgs, multiErrorSeparator)
}

// ErrResult returns true if not errors have been added
func (me MultiError) ErrResult() error {
	if len(me.errors) == 0 {
		return nil
	}
	return &me
}

// Unwrap returns the underlying error
func (me MultiError) Unwrap() error {
	// Avoid infinite loop if a MultiError instance is used to call `errors.Is` function.
	// See https://github.com/golang/go/blob/master/src/errors/wrap.go#L55
	return nil
}

// Add adds a new error
func (me *MultiError) Add(err error) {
	me.errors = append(me.errors, err)
}

// Is provides the `errors.Is` feature
func (me MultiError) Is(err error) bool {
	for _, e := range me.errors {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

// Errors return all underlying errors
func (me MultiError) Errors() []error {
	return me.errors
}
