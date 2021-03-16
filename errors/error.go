package gierrors

import "fmt"

// Err holds a description of an error along with information about
// where the error was created.
//
// It may be embedded in custom error types to add extra information that
// this errors package can understand.
type Err struct {
	// message holds an annotation of the error.
	message string

	// cause holds the cause of the error as returned
	// by the Cause method.
	cause error

	// previous holds the previous error in the error stack, if any.
	previous error

	// file and line hold the source code location where the error was
	// created.
	file string
	line int
}

// NewErr is used to return an Err for the purpose of embedding in other
// structures.  The location is not specified, and needs to be set with a call
// to SetLocation.
//
// For example:
//     type FooError struct {
//         errors.Err
//         code int
//     }
//
//     func NewFooError(code int) error {
//         err := &FooError{errors.NewErr("foo"), code}
//         err.SetLocation(1)
//         return err
//     }
func NewErr(format string, args ...interface{}) Err {
	return Err{
		message: fmt.Sprintf(format, args...),
	}
}

// NewErrWithCause is used to return an Err with cause by other error for the purpose of embedding in other
// structures. The location is not specified, and needs to be set with a call
// to SetLocation.
//
// For example:
//     type FooError struct {
//         errors.Err
//         code int
//     }
//
//     func (e *FooError) Annotate(format string, args ...interface{}) error {
//         err := &FooError{errors.NewErrWithCause(e.Err, format, args...), e.code}
//         err.SetLocation(1)
//         return err
//     })
func NewErrWithCause(other error, format string, args ...interface{}) Err {
	return Err{
		message:  fmt.Sprintf(format, args...),
		cause:    Cause(other),
		previous: other,
	}
}

// Location is the file and line of where the error was most recently
// created or annotated.
func (e *Err) Location() (filename string, line int) {
	return e.file, e.line
}

// Underlying returns the previous error in the error stack, if any. A client
// should not ever really call this method.  It is used to build the error
// stack and should not be introspected by client calls.  Or more
// specifically, clients should not depend on anything but the `Cause` of an
// error.
func (e *Err) Underlying() error {
	return e.previous
}

// Cause returns the most recent error in the error stack that
// meets one of these criteria: the original error that was raised; the new
// error that was passed into the Wrap function; the most recently masked
// error; or nil if the error itself is considered the Cause.  Normally this
// method is not invoked directly, but instead through the Cause stand alone
// function.
func (e *Err) Cause() error {
	return e.cause
}

// Message returns the message stored with the most recent location. This is
// the empty string if the most recent call was Trace, or the message stored
// with Annotate or Mask.
func (e *Err) Message() string {
	return e.message
}

// Error implements error.Error.
func (e *Err) Error() string {
	// We want to walk up the stack of errors showing the annotations
	// as long as the cause is the same.
	err := e.previous
	if !sameError(Cause(err), e.cause) && e.cause != nil {
		err = e.cause
	}
	switch {
	case err == nil:
		return e.message
	case e.message == "":
		return err.Error()
	}
	return fmt.Sprintf("%s: %v", e.message, err)
}

// Format implements fmt.Formatter
// When printing errors with %+v it also prints the stack trace.
// %#v unsurprisingly will print the real underlying type.
func (e *Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			fmt.Fprintf(s, "%s", ErrorStack(e))
			return
		case s.Flag('#'):
			// avoid infinite recursion by wrapping e into a type
			// that doesn't implement Formatter.
			fmt.Fprintf(s, "%#v", (*unformatter)(e))
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	default:
		fmt.Fprintf(s, "%%!%c(%T=%s)", verb, e, e.Error())
	}
}
