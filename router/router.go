package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	// Initialize Routes
	iniatializeRoutes(router)
	router.Run(":8080")
}
