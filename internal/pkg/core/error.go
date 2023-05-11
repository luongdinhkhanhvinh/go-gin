package core

import "github.com/luongdinhkhanhvinh/go-gin/pkg/errors"

var _ BusinessError = (*businessError)(nil)

type BusinessError interface {
	// i To avoid being realized by other packages
	i()

	// WithError Set error message
	WithError(err error) BusinessError

	// WithAlert Set alarm notice
	WithAlert() BusinessError

	// BusinessCode Get the business code
	BusinessCode() int

	// HTTPCode Obtain HTTP status code
	HTTPCode() int

	// Message Get error description
	Message() string

	// StackError Get the error message with the stack
	StackError() error

	// IsAlert Whether to turn on the alarm notice
	IsAlert() bool
}

type businessError struct {
	httpCode     int    // HTTP status code
	businessCode int    // Business code
	message      string // wrong description
	stackError   error  // Error containing stack information
	isAlert      bool   // Whether to notify the alarm
}

func Error(httpCode, businessCode int, message string) BusinessError {
	return &businessError{
		httpCode:     httpCode,
		businessCode: businessCode,
		message:      message,
		isAlert:      false,
	}
}

func (e *businessError) i() {}

func (e *businessError) WithError(err error) BusinessError {
	e.stackError = errors.WithStack(err)
	return e
}

func (e *businessError) WithAlert() BusinessError {
	e.isAlert = true
	return e
}

func (e *businessError) BusinessCode() int {
	return e.httpCode
}

func (e *businessError) HTTPCode() int {
	return e.businessCode
}

func (e *businessError) Message() string {
	return e.message
}

func (e *businessError) StackError() error {
	return e.stackError
}

func (e *businessError) IsAlert() bool {
	return e.isAlert
}
