package errors

// serviceUnavailable represents an error when an operation is serviceUnavailable.
type serviceUnavailable struct {
	Err
}

// ServiceUnavailablef returns an error which satisfies IsServiceUnavailable().
func ServiceUnavailablef(format string, args ...interface{}) error {
	return &serviceUnavailable{wrap(nil, format, "", args...)}
}

// NewServiceUnavailable returns an error which wraps err and satisfies
// IsServiceUnavailable().
func NewServiceUnavailable(err error, msg string) error {
	return &serviceUnavailable{wrap(err, msg, "")}
}

// IsServiceUnavailable reports whether err was created with ServiceUnavailablef() or
// NewServiceUnavailable().
func IsServiceUnavailable(err error) bool {
	err = Cause(err)
	_, ok := err.(*serviceUnavailable)
	return ok
}
