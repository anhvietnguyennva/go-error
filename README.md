# Go Error Lib

## Overview
This is the library that defines errors and error handlers for Go application

## Add this lib to your project
- Step 1: 
```
$ export GOPRIVATE=github.com/anhvietnguyennva/go-error
```
- Step 2: Add file `tool/tool.go` with content:
```
package tool

import (
	_ "github.com/anhvietnguyennva/go-error/tool"
)
```
- Step 3: 
```
$ go mod tidy
$ go mod vendor
```

## Update to latest version
```
$ go get -u github.com/anhvietnguyennva/go-error
$ go mod vendor
```

## How to use

### `Error`
These errors should be used in the internal application

### `RestAPIError`
These errors should be used in the application interface layer (API) of your service

### Transforms Error to RestAPIError
```
package main

import (
	"fmt"

	"github.com/anhvietnguyennva/go-error/pkg/error"
	t "github.com/anhvietnguyennva/go-error/pkg/transformer"
)

func main() {
	err := error.NewErrNotFound(nil)
	transformer := t.RestTransformerInstance()
	apiErr := transformer.ErrToRestAPIErr(err)
	fmt.Println(apiErr.Error())
}

```

### Register new Error and new RestAPIError
- There are functions can be used to create customized errors. For `Error`, it's `error.NewError(code string, message string, entities []string, rootCause error)` and it's `errors.NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error)` for `RestAPIError`.
- You can use these above functions to create your customized errors directly. But you should define new constructors for your custom errors, so it can be reused. For example:
```
package main

import (
	"net/http"

	e "github.com/anhvietnguyennva/go-error/pkg/error"
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

```
- After defining your custom errors, you have to register the function used to transform your custom `Error` to your custom `RestAPIError`. For example:
```
package main

import (
	"net/http"

	e "github.com/anhvietnguyennva/go-error/pkg/error"
	"github.com/anhvietnguyennva/go-error/pkg/transformer"
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
```
- NOTE: 
  - Each error should have a unique error code. Otherwise, it can lead to unexpected results when transforming. So You should not define your custom error code as one of the predefined error codes in "go-error".
  - The `ClientErrorCode` should contain information about HTTP Status. It makes the error code more meaningful

### For gin framework
- This lib provides the function `ValidationErrToRestAPIErr(err error)` which can be used when binding and validating the request.
