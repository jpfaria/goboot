package errors

type causer interface {
	Cause() error
}
