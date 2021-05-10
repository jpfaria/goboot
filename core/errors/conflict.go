package errors

// conflict represents an error when something is not supported.
type conflict struct {
	Err
}

// Conflictf returns an error which satisfies IsConflict().
func Conflictf(format string, args ...interface{}) error {
	return &conflict{wrap(nil, format, "", args...)}
}

// NewConflict returns an error which wraps err and satisfies
// IsConflict().
func NewConflict(err error, msg string) error {
	return &conflict{wrap(err, msg, "")}
}

// IsConflict reports whether the error was created with
// Conflictf() or NewConflict().
func IsConflict(err error) bool {
	err = Cause(err)
	_, ok := err.(*conflict)
	return ok
}
