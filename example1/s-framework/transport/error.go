package transport

import (
	"errors"
	"fmt"
)

type Error struct {
	Code  string
	Cause error
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, cause: %s", e.Code, e.Cause.Error())
}
func (e Error) WithCause(err error) Error {
	e.Cause = err
	return e
}

func NewError(code string) Error {
	return Error{Code: code}
}

const (
	ErrCodeUnknownMessageType = "unknown_message_type"
	ErrCodeInternal           = "internal"
	ErrCodeTimeout            = "timeout"
)

var (
	ErrUnknownRecipient = errors.New("unknown recipient")
)
