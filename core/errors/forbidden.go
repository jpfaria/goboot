package errors

// forbidden represents an error when a request cannot be completed because of
// missing privileges
type forbidden struct {
	Err
}

// Forbiddenf returns an error which satistifes IsForbidden()
func Forbiddenf(format string, args ...interface{}) error {
	return &forbidden{wrap(nil, format, "", args...)}
}

// NewForbidden returns an error which wraps err that satisfies
// IsForbidden().
func NewForbidden(err error, msg string) error {
	return &forbidden{wrap(err, msg, "")}
}

// IsForbidden reports whether err was created with Forbiddenf() or
// NewForbidden().
func IsForbidden(err error) bool {
	err = Cause(err)
	_, ok := err.(*forbidden)
	return ok
}
