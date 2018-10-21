package apierror

import "github.com/pkg/errors"

const (
	CodeInternal   = "INTERNAL"
	CodeBadRequest = "BAD_REQUEST"
)

var (
	BadRequest = Error{Code: CodeBadRequest}
	Internal   = Error{Code: CodeInternal}
)

type Error struct {
	Cause error
	Code  string
}

func (e Error) Error() string {
	return ""
}

func (e Error) WithCause(err error) Error {
	e.Cause = errors.Wrap(e.Cause, err.Error())
	return e
}
