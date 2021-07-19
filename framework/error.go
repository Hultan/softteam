package framework

import "fmt"

type Error struct {
}

// CheckWithPanic : panics on error
func (e *Error) CheckWithPanic(err error, message string) {
	if err != nil {
		panic(err.Error() + " : " + message)
	}
}

// CheckWithoutPanic : panics on error
func (e* Error) CheckWithoutPanic(err error, message string) bool {
	if err != nil {
		fmt.Println(err.Error() + " : " + message)
		return true
	}
	return false
}
