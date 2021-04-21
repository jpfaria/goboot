package errors

import (
	"reflect"
	"runtime"
)

// helper for Format
type unformatter Err

func (unformatter) Format() { /* break the fmt.Formatter interface */ }

// SetLocation records the source location of the error at callDepth stack
// frames above the call.
func (e *Err) SetLocation(callDepth int) {
	_, file, line, _ := runtime.Caller(callDepth + 1)
	e.file = trimSourcePath(file)
	e.line = line
}

// StackTrace returns one string for each location recorded in the stack of
// errors. The first value is the originating error, with a line for each
// other annotation or tracing of the error.
func (e *Err) StackTrace() []string {
	return errorStack(e)
}

// Ideally we'd have a way to check identity, but deep equals will do.
func sameError(e1, e2 error) bool {
	return reflect.DeepEqual(e1, e2)
}
