package transformer

import (
	"fmt"

	"github.com/anhvietnguyennva/go-error/pkg/constant"
	e "github.com/anhvietnguyennva/go-error/pkg/error"
)

type IErrorTransformer interface {
	RestAPIErrToErr(restAPIErr *e.RestAPIError) *e.Error
}

type errTransformFunc func(rootCause error, entities ...string) *e.Error

type errTransformer struct {
	mapping map[int]errTransformFunc
}

var errTransformerInstance *errTransformer

func initErrTransformerInstance() {
	if errTransformerInstance == nil {
		errTransformerInstance = &errTransformer{
			mapping: make(map[int]errTransformFunc),
		}

		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeRequired, e.NewErrRequired)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeNotAcceptedValue, e.NewErrNotAcceptedValue)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeOutOfRange, e.NewErrOutOfRange)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeInvalidFormat, e.NewErrInvalidFormat)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeInvalid, e.NewErrInvalid)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeUnauthenticated, e.NewErrUnauthenticated)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeUnauthorized, e.NewErrUnauthorized)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeNotFound, e.NewErrNotFound)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeDuplicate, e.NewErrDuplicate)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeAlreadyExists, e.NewErrAlreadyExits)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeTooManyRequests, e.NewErrTooManyRequests)
		errTransformerInstance.RegisterTransformFunc(constant.ClientErrCodeInternal, e.NewErrUnknown)
	}
}

func ErrTransformerInstance() IErrorTransformer {
	return errTransformerInstance
}

// RestAPIErrToErr transforms RestAPIError to Error
func (t *errTransformer) RestAPIErrToErr(restAPIErr *e.RestAPIError) *e.Error {
	f := t.mapping[restAPIErr.Code]
	if f == nil {
		return e.NewErrUnknown(fmt.Errorf("can not transform rest API error, error: %v", restAPIErr))
	}
	return f(restAPIErr, restAPIErr.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform RestAPIError to Error
// if the restAPIErrorCode is already registered, the old transform function will be overridden
func (t *errTransformer) RegisterTransformFunc(restAPIErrCode int, transformFunc errTransformFunc) {
	t.mapping[restAPIErrCode] = transformFunc
}
