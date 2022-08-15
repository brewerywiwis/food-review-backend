package rest

import (
	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
	"github.com/gin-gonic/gin"
)

// Server represents rest server interface
type Server interface {
	Run(addr string) error
}

// GinServer is implementation of rest server by Gin
type GinServer struct {
	engine *gin.Engine
}

// Run is used to start a gin server
func (s *GinServer) Run(addr string) error {
	return s.engine.Run(addr)
}

// NewServer returns a new Server
func NewServer(foodService food.Service) Server {
	server := &GinServer{gin.Default()}
	registerRoutes(server, foodService)

	return server
}
