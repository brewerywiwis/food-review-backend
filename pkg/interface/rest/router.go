package rest

import "github.com/gin-gonic/gin"

func registerV1Route(router *gin.RouterGroup) {
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "success"})
	})
}

func registerRoutes(server *GinServer) {
	api := server.engine.Group("/api")
	v1 := api.Group("/v1")
	registerV1Route(v1)
}
