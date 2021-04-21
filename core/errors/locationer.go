package errors

type locationer interface {
	Location() (string, int)
}
