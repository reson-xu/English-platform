package user

import apperrors "github.com/reson-xu/english-platform/internal/platform/errors"

var (
	ErrEmailAlreadyExists = apperrors.New(apperrors.CodeConflict, "email already exists")
	ErrInvalidCredentials = apperrors.New(apperrors.CodeUnauthorized, "email or password is incorrect")
	ErrUserNotFound       = apperrors.New(apperrors.CodeNotFound, "user not found")
	ErrUserDisabled       = apperrors.New(apperrors.CodeForbidden, "user is disabled")
)

func invalidArgument(message string) *apperrors.AppError {
	return apperrors.New(apperrors.CodeInvalidArgument, message)
}
