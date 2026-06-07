package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reson-xu/english-platform/internal/adapters/httpapi"
	userdto "github.com/reson-xu/english-platform/internal/models/dto/user"
	user "github.com/reson-xu/english-platform/internal/modules/user"
	apperrors "github.com/reson-xu/english-platform/internal/platform/errors"
)

type UserHandler struct {
	users *user.Service
}

func NewUserHandler(users *user.Service) *UserHandler {
	return &UserHandler{users: users}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req userdto.RegisterReq
	if err := httpapi.DecodeJSON(c, &req); err != nil {
		httpapi.WriteError(c, apperrors.ErrInvalidJSON)
		return
	}

	resp, err := h.users.Register(c.Request.Context(), req)
	if err != nil {
		httpapi.WriteError(c, err)
		return
	}

	httpapi.WriteData(c, http.StatusCreated, resp)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req userdto.LoginReq
	if err := httpapi.DecodeJSON(c, &req); err != nil {
		httpapi.WriteError(c, apperrors.ErrInvalidJSON)
		return
	}

	resp, err := h.users.Login(c.Request.Context(), req)
	if err != nil {
		httpapi.WriteError(c, err)
		return
	}

	httpapi.WriteData(c, http.StatusOK, resp)
}
