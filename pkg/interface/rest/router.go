package rest

import (
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/gin-gonic/gin"
)

func registerV1Routes(router *gin.RouterGroup, foodService food.Service) {
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "success"})
	})

	foodController := foodController{foodService}
	router.GET("/foods", foodController.listAllFood)
	router.POST("/food", foodController.createFood)
}

func registerRoutes(server *GinServer, foodService food.Service) {
	api := server.engine.Group("/api")
	v1 := api.Group("/v1")
	registerV1Routes(v1, foodService)
}
