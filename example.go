package main

import (
	"net/http"

	e "github.com/telifi/go-error/pkg/error"
	"github.com/telifi/go-error/pkg/transformer"
)

const (
	ErrCodeCustomized = "CUSTOMIZED"
	ErrMsgCustomized  = "Customized error"

	ClientErrCodeCustomized = 40099
	ClientErrMsgCustomized  = "Customized client error"
)

func NewErrCustomized(rootCause error, entities ...string) *e.Error {
	return e.NewError(ErrCodeCustomized, ErrMsgCustomized, entities, rootCause)
}

func NewRestAPIErrCustomized(rootCause error, entities ...string) *e.RestAPIError {
	message := e.AppendEntitiesToErrMsg(ClientErrMsgCustomized, entities)
	return e.NewRestAPIError(http.StatusBadRequest, ClientErrCodeCustomized, message, entities, rootCause)
}

func main() {
	t := transformer.RestTransformerInstance()
	t.RegisterTransformFunc(ErrCodeCustomized, NewRestAPIErrCustomized)
}
