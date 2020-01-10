package myclient

import (
	"errors"
)

var (
	ErrValidKey    = errors.New("key is invalid")
	ErrValidNumber = errors.New("number is invalid")
	ErrKeyNotExist = errors.New("key is not existing")
)

func NewMyError(text string) *MyError {
	return &MyError{
		Text: text,
	}
}

type MyError struct {
	Inner error
	Text  string
}

func (e MyError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.Text != "" {
		return e.Text
	} else {
		return "invalid error!!"
	}
}
