package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const HeaderRequestID = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = newRequestID()
		}

		c.Header(HeaderRequestID, requestID)
		c.Set(HeaderRequestID, requestID)
		c.Next()
	}
}

func RequestIDFromContext(c *gin.Context) string {
	value, ok := c.Get(HeaderRequestID)
	if !ok {
		return ""
	}

	requestID, ok := value.(string)
	if !ok {
		return ""
	}

	return requestID
}

func newRequestID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
