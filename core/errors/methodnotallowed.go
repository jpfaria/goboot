package errors

// methodNotAllowed represents an error when an HTTP request
// is made with an inappropriate method.
type methodNotAllowed struct {
	Err
}

// MethodNotAllowedf returns an error which satisfies IsMethodNotAllowed().
func MethodNotAllowedf(format string, args ...interface{}) error {
	return &methodNotAllowed{wrap(nil, format, "", args...)}
}

// NewMethodNotAllowed returns an error which wraps err that satisfies
// IsMethodNotAllowed().
func NewMethodNotAllowed(err error, msg string) error {
	return &methodNotAllowed{wrap(err, msg, "")}
}

// IsMethodNotAllowed reports whether err was created with MethodNotAllowedf() or
// NewMethodNotAllowed().
func IsMethodNotAllowed(err error) bool {
	err = Cause(err)
	_, ok := err.(*methodNotAllowed)
	return ok
}
