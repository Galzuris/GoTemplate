package general

import (
	timeHandler "backend/internal/domain/general/time/handler"

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

	handler.GET("/time", timeHandler.GetTime)

	// TEST HANDLER
	handler.GET("/abort", func(ctx *gin.Context) {
		ctx.AbortWithError(http.StatusForbidden, errors.New("forbidden"))
	})

	handler.GET("/exception", func(ctx *gin.Context) {
		ctx.Error(exception.NewApiException(http.StatusUnprocessableEntity, "user not accepted"))
	})
}
