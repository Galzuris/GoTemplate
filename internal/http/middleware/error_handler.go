package middleware

import (
	"backend/internal/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			if er, ok := err.Err.(*exception.ApiException); ok {
				c.AbortWithStatusJSON(er.Status, gin.H{"message": er.Message})
				return
			}
		}

		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
	}
}
