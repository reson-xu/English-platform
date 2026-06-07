package httpapi

import (
	"github.com/gin-gonic/gin"
	"github.com/reson-xu/english-platform/internal/models/dto"
	apperrors "github.com/reson-xu/english-platform/internal/platform/errors"
)

func DecodeJSON(c *gin.Context, out any) error {
	return c.ShouldBindJSON(out)
}

func WriteData(c *gin.Context, status int, value any) {
	c.JSON(status, dto.DataResp{Data: value})
}

func WriteError(c *gin.Context, err error) {
	appErr := apperrors.From(err)
	if appErr == nil {
		appErr = apperrors.ErrInternal
	}

	c.JSON(apperrors.HTTPStatus(appErr.Code), dto.ErrorEnvelopeResp{
		Error: dto.ErrorResp{
			Code:    string(appErr.Code),
			Message: appErr.Message,
		},
	})
}
