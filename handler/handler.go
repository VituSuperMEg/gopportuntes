package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Todas",
	})
}
func CreateOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post",
	})
}
func ShowOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get",
	})
}
func DeleteOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete",
	})
}
func UpdateOpeningHanlder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update",
	})
}
