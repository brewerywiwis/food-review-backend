package rest

import "github.com/gin-gonic/gin"

func registerRoutes(server *gin.Engine) {
	api := server.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "success"})
	})

}
