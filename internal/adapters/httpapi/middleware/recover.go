package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/reson-xu/english-platform/internal/models/dto"
	apperrors "github.com/reson-xu/english-platform/internal/platform/errors"
)

func Recover(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				logger.Error("http panic recovered",
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
					"request_id", RequestIDFromContext(c),
					"panic", recovered,
				)
				writeRecoveryError(c)
			}
		}()

		c.Next()
	}
}

func writeRecoveryError(c *gin.Context) {
	appErr := apperrors.ErrInternal
	c.AbortWithStatusJSON(apperrors.HTTPStatus(appErr.Code), dto.ErrorEnvelopeResp{
		Error: dto.ErrorResp{
			Code:    string(appErr.Code),
			Message: appErr.Message,
		},
	})
}
