package error

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anhvietnguyennva/go-error/pkg/constant"
)

type RestAPIError struct {
	HttpStatus    int           `json:"-"`
	Code          int           `json:"code"`
	Message       string        `json:"message"`
	ErrorEntities []string      `json:"errorEntities"`
	Details       []interface{} `json:"details"`
	RootCause     error         `json:"-"`
}

func NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error) *RestAPIError {
	return &RestAPIError{
		HttpStatus:    httpStatus,
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *RestAPIError) Error() string {
	return fmt.Sprintf("API ERROR: {Code: %d, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func AppendEntitiesToErrMsg(message string, entities []string) string {
	if len(entities) > 0 {
		message += ": "
		message += strings.Join(entities, ",")
	}
	return message
}

func NewRestAPIErrRequired(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgRequired, entities)
	return NewRestAPIError(http.StatusBadRequest, constant.ClientErrCodeRequired, message, entities, rootCause)
}

func NewRestAPIErrInvalidFormat(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgInvalidFormat, entities)
	return NewRestAPIError(http.StatusBadRequest, constant.ClientErrCodeInvalidFormat, message, entities, rootCause)
}

func NewRestAPIErrInvalid(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgInvalid, entities)
	return NewRestAPIError(http.StatusBadRequest, constant.ClientErrCodeInvalid, message, entities, rootCause)
}

func NewRestAPIErrNotAcceptedValue(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgNotAcceptedValue, entities)
	return NewRestAPIError(http.StatusBadRequest, constant.ClientErrCodeNotAcceptedValue, message, entities, rootCause)
}

func NewRestAPIErrOutOfRange(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgOutOfRange, entities)
	return NewRestAPIError(http.StatusBadRequest, constant.ClientErrCodeOutOfRange, message, entities, rootCause)
}

func NewRestAPIErrUnauthenticated(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgUnauthenticated, entities)
	return NewRestAPIError(http.StatusUnauthorized, constant.ClientErrCodeUnauthenticated, message, entities, rootCause)
}

func NewRestAPIErrUnauthorized(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgUnauthorized, entities)
	return NewRestAPIError(http.StatusForbidden, constant.ClientErrCodeUnauthorized, message, entities, rootCause)
}

func NewRestAPIErrNotFound(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgNotFound, entities)
	return NewRestAPIError(http.StatusNotFound, constant.ClientErrCodeNotFound, message, entities, rootCause)
}

func NewRestAPIErrDuplicate(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgDuplicate, entities)
	return NewRestAPIError(http.StatusConflict, constant.ClientErrCodeDuplicate, message, entities, rootCause)
}

func NewRestAPIErrAlreadyExits(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgAlreadyExists, entities)
	return NewRestAPIError(http.StatusConflict, constant.ClientErrCodeAlreadyExists, message, entities, rootCause)
}

func NewRestAPIErrTooManyRequests(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(constant.ClientErrMsgTooManyRequests, entities)
	return NewRestAPIError(http.StatusTooManyRequests, constant.ClientErrCodeTooManyRequests, message, entities, rootCause)
}

func NewRestAPIErrInternal(rootCause error, entities ...string) *RestAPIError {
	return NewRestAPIError(http.StatusInternalServerError, constant.ClientErrCodeInternal, constant.ClientErrMsgInternal, entities, rootCause)
}
