package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reson-xu/english-platform/internal/adapters/httpapi/middleware"
	"github.com/reson-xu/english-platform/internal/models/dto"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func NewRouter(userHandler UserHandler, logger *slog.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middleware.Recover(logger),
		middleware.RequestID(),
		middleware.Logging(logger),
		cors(),
	)

	router.GET("/api/health", handleHealth)

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	return router
}

func handleHealth(c *gin.Context) {
	WriteData(c, http.StatusOK, dto.HealthResp{
		AppName: "English Platform",
		Status:  "ok",
	})
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
