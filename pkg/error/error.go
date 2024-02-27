package error

import (
	"fmt"

	"github.com/telifi/go-error/pkg/constant"
)

type Error struct {
	Code          string   `json:"code"`
	Message       string   `json:"message"`
	ErrorEntities []string `json:"errorEntities"`
	RootCause     error    `json:"-"`
}

func NewError(code string, message string, entities []string, rootCause error) *Error {
	return &Error{
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("ERROR: {Code: %s, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func NewErrRequired(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeRequired, constant.ErrMsgRequired, entities, rootCause)
}

func NewErrInvalidFormat(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeInvalidFormat, constant.ErrMsgInvalidFormat, entities, rootCause)
}

func NewErrInvalid(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeInvalid, constant.ErrMsgInvalid, entities, rootCause)
}

func NewErrNotAcceptedValue(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeNotAcceptedValue, constant.ErrMsgNotAcceptedValue, entities, rootCause)
}

func NewErrOutOfRange(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeOutOfRange, constant.ErrMsgOutOfRange, entities, rootCause)
}

func NewErrUnauthenticated(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeUnauthenticated, constant.ErrMsgUnauthenticated, entities, rootCause)
}

func NewErrUnauthorized(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeUnauthorized, constant.ErrMsgUnauthorized, entities, rootCause)
}

func NewErrNotFound(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeNotFound, constant.ErrMsgNotFound, entities, rootCause)
}

func NewErrDuplicate(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeDuplicate, constant.ErrMsgDuplicate, entities, rootCause)
}

func NewErrAlreadyExits(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeAlreadyExists, constant.ErrMsgAlreadyExists, entities, rootCause)
}

func NewErrTooManyRequests(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeTooManyRequests, constant.ErrMsgTooManyRequests, entities, rootCause)
}

func NewErrUnknown(rootCause error, entities ...string) *Error {
	return NewError(constant.ErrCodeUnknown, constant.ErrMsgUnknown, entities, rootCause)
}
