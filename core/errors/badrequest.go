package errors

// badRequest represents an error when a request has bad parameters.
type badRequest struct {
	Err
}

// BadRequestf returns an error which satisfies IsBadRequest().
func BadRequestf(format string, args ...interface{}) error {
	return &badRequest{wrap(nil, format, "", args...)}
}

// NewBadRequest returns an error which wraps err that satisfies
// IsBadRequest().
func NewBadRequest(err error, msg string) error {
	return &badRequest{wrap(err, msg, "")}
}

// IsBadRequest reports whether err was created with BadRequestf() or
// NewBadRequest().
func IsBadRequest(err error) bool {
	err = Cause(err)
	_, ok := err.(*badRequest)
	return ok
}
