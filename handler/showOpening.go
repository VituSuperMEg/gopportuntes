package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get",
	})
}
