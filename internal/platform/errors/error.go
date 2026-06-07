package errors

import stderrors "errors"

var (
	ErrInvalidJSON = New(CodeInvalidArgument, "request body must be valid JSON")
	ErrInternal    = New(CodeInternal, "internal server error")
)

type AppError struct {
	Code    Code
	Message string
	cause   error
}

func New(code Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(cause error, code Code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		cause:   cause,
	}
}

func From(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if stderrors.As(err, &appErr) {
		return appErr
	}

	return ErrInternal
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.cause != nil {
		return e.cause.Error()
	}

	return e.Message
}

func (e *AppError) Is(target error) bool {
	targetErr, ok := target.(*AppError)
	if !ok {
		return false
	}

	return e.Code == targetErr.Code
}

func (e *AppError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.cause
}
