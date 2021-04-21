package errors

// notImplemented represents an error when something is not
// implemented.
type notImplemented struct {
	Err
}

// NotImplementedf returns an error which satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	return &notImplemented{wrap(nil, format, " not implemented", args...)}
}

// NewNotImplemented returns an error which wraps err and satisfies
// IsNotImplemented().
func NewNotImplemented(err error, msg string) error {
	return &notImplemented{wrap(err, msg, "")}
}

// IsNotImplemented reports whether err was created with
// NotImplementedf() or NewNotImplemented().
func IsNotImplemented(err error) bool {
	err = Cause(err)
	_, ok := err.(*notImplemented)
	return ok
}
