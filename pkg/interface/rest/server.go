package rest

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	server := gin.Default()

	registerRoutes(server)

	return server
}
