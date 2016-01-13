package main

// IError interface
type IError interface {
	Error() string
}

// Error represents the error
type Error struct {
	errString string
}

// Error returns the error message of the object
func (e *Error) Error() string {
	return e.errString
}
