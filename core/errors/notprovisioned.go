package errors

// notProvisioned represents an error when something is not yet provisioned.
type notProvisioned struct {
	Err
}

// NotProvisionedf returns an error which satisfies IsNotProvisioned().
func NotProvisionedf(format string, args ...interface{}) error {
	return &notProvisioned{wrap(nil, format, " not provisioned", args...)}
}

// NewNotProvisioned returns an error which wraps err that satisfies
// IsNotProvisioned().
func NewNotProvisioned(err error, msg string) error {
	return &notProvisioned{wrap(err, msg, "")}
}

// IsNotProvisioned reports whether err was created with NotProvisionedf() or
// NewNotProvisioned().
func IsNotProvisioned(err error) bool {
	err = Cause(err)
	_, ok := err.(*notProvisioned)
	return ok
}
