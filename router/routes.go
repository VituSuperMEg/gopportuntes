package router

import (
	"github.com/VituSuperMEg/gopportuntes/handler"
	"github.com/gin-gonic/gin"
)

func iniatializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningHanlder)
		v1.POST("/opening", handler.CreateOpeningHanlder)
		v1.PUT("/opening", handler.UpdateOpeningHanlder)
		v1.DELETE("/opening", handler.DeleteOpeningHanlder)
		v1.GET("/opening", handler.ShowOpeningHanlder)
	}
}
