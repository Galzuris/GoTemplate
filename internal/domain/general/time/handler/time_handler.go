package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTime(ctx *gin.Context) {
	now := time.Now().UTC()
	formatted := now.Format("2006-01-02T15:04:05-0700")

	ctx.JSON(http.StatusOK, gin.H{"time": formatted})
}
