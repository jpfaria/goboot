package errors

// timeout represents an error on timeout.
type timeout struct {
	Err
}

// Timeoutf returns an error which satisfies IsTimeout().
func Timeoutf(format string, args ...interface{}) error {
	return &timeout{wrap(nil, format, " timeout", args...)}
}

// NewTimeout returns an error which wraps err that satisfies
// IsTimeout().
func NewTimeout(err error, msg string) error {
	return &timeout{wrap(err, msg, "")}
}

// IsTimeout reports whether err was created with Timeoutf() or
// NewTimeout().
func IsTimeout(err error) bool {
	err = Cause(err)
	_, ok := err.(*timeout)
	return ok
}
