package errors

// internal represents an error when a request cannot be completed because of
// missing privileges
type internal struct {
	Err
}

// Internalf returns an error which satistifes IsInternal()
func Internalf(format string, args ...interface{}) error {
	return &internal{wrap(nil, format, "", args...)}
}

// NewInternal returns an error which wraps err that satisfies
// IsInternal().
func NewInternal(err error, msg string) error {
	return &internal{wrap(err, msg, "")}
}

// IsInternal reports whether err was created with Internalf() or
// NewInternal().
func IsInternal(err error) bool {
	err = Cause(err)
	_, ok := err.(*internal)
	return ok
}
