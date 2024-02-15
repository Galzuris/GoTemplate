package general

import (
	"backend/internal/exception"
	"backend/internal/http/middleware"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(middleware.ErrorHandler())

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	handler.GET("/abort", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			exception.NewApiException(http.StatusUnprocessableEntity, "user not found"),
		)
	})

	handler.GET("/exception", func(ctx *gin.Context) {
		ctx.Error(errors.New("user not found"))
	})
}
