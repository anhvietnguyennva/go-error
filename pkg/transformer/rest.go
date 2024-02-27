package transformer

import (
	"encoding/json"
	errs "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/telifi/go-error/pkg/constant"
	e "github.com/telifi/go-error/pkg/error"
)

type IRestTransformer interface {
	ErrToRestAPIErr(err *e.Error) *e.RestAPIError
	ValidationErrToRestAPIErr(err error) *e.RestAPIError
	RegisterTransformFunc(errCode string, transformFunc restTransformFunc)
	RegisterValidationTag(tag string, function restTransformFunc)
}

type restTransformFunc func(rootCause error, entities ...string) *e.RestAPIError

type restTransformer struct {
	mapping       map[string]restTransformFunc
	validationErr map[string]restTransformFunc
}

var restTransformerInstance *restTransformer

func initRestTransformerInstance() {
	if restTransformerInstance == nil {
		restTransformerInstance = &restTransformer{
			mapping:       make(map[string]restTransformFunc),
			validationErr: make(map[string]restTransformFunc),
		}

		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeRequired, e.NewRestAPIErrRequired)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeNotAcceptedValue, e.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeOutOfRange, e.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeInvalidFormat, e.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeInvalid, e.NewRestAPIErrInvalid)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeUnauthenticated, e.NewRestAPIErrUnauthenticated)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeUnauthorized, e.NewRestAPIErrUnauthorized)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeNotFound, e.NewRestAPIErrNotFound)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeDuplicate, e.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeAlreadyExists, e.NewRestAPIErrAlreadyExits)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeTooManyRequests, e.NewRestAPIErrTooManyRequests)
		restTransformerInstance.RegisterTransformFunc(constant.ErrCodeUnknown, e.NewRestAPIErrInternal)

		restTransformerInstance.RegisterValidationTag("required", e.NewRestAPIErrRequired)
		restTransformerInstance.RegisterValidationTag("oneof", e.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterValidationTag("min", e.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("max", e.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("numeric", e.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("unique", e.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterValidationTag("hexadecimal", e.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("email", e.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("url", e.NewRestAPIErrInvalidFormat)
	}
}

func RestTransformerInstance() IRestTransformer {
	return restTransformerInstance
}

// ValidationErrToRestAPIErr transforms ValidationError to RestAPIError
// this function will be used when bind JSON request to DTO in gin framework
func (t *restTransformer) ValidationErrToRestAPIErr(err error) *e.RestAPIError {
	var validationErrs validator.ValidationErrors
	var unmarshalTypeErr *json.UnmarshalTypeError
	var jsonSynTaxErr *json.SyntaxError
	var numErr *strconv.NumError
	if errs.As(err, &validationErrs) {
		validationErr := validationErrs[0]
		return t.apiErrForTag(validationErr.Tag(), err, validationErr.Field())
	}
	if errs.As(err, &unmarshalTypeErr) {
		field := unmarshalTypeErr.Field
		fieldArr := strings.Split(field, ".")
		return e.NewRestAPIErrInvalidFormat(err, fieldArr[len(fieldArr)-1])
	}
	if errs.As(err, &jsonSynTaxErr) {
		return e.NewRestAPIErrInvalidFormat(err)
	}
	if errs.As(err, &numErr) {
		return e.NewRestAPIErrInvalidFormat(err)
	}
	return e.NewRestAPIErrInternal(err)
}

// ErrToRestAPIErr transforms Error to RestAPIError
func (t *restTransformer) ErrToRestAPIErr(err *e.Error) *e.RestAPIError {
	f := t.mapping[err.Code]
	if f == nil {
		return e.NewRestAPIErrInternal(fmt.Errorf("can not transform error, error: %v", err))
	}
	return f(err, err.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform DomainError to RestAPIError
// if the domainErrCode is already registered, the old transform function will be overridden
func (t *restTransformer) RegisterTransformFunc(domainErrCode string, function restTransformFunc) {
	t.mapping[domainErrCode] = function
}

// RegisterValidationTag is used to define new validation tag and respective API error
// if the validation tag is already registered, the old respective API error will be overridden
func (t *restTransformer) RegisterValidationTag(tag string, function restTransformFunc) {
	t.validationErr[tag] = function
}

// apiErrForTag return RestAPIError which corresponds to the validation tag
func (t *restTransformer) apiErrForTag(tag string, err error, fields ...string) *e.RestAPIError {
	f := t.validationErr[tag]
	if f == nil {
		return e.NewRestAPIErrInternal(err)
	}
	return f(err, fields...)
}
