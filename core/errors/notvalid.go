package errors

// notValid represents an error when something is not valid.
type notValid struct {
	Err
}

// NotValidf returns an error which satisfies IsNotValid().
func NotValidf(format string, args ...interface{}) error {
	return &notValid{wrap(nil, format, " not valid", args...)}
}

// NewNotValid returns an error which wraps err and satisfies IsNotValid().
func NewNotValid(err error, msg string) error {
	return &notValid{wrap(err, msg, "")}
}

// IsNotValid reports whether the error was created with NotValidf() or
// NewNotValid().
func IsNotValid(err error) bool {
	err = Cause(err)
	_, ok := err.(*notValid)
	return ok
}
