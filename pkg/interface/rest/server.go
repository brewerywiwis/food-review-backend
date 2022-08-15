package rest

import (
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
func NewServer() Server {
	server := &GinServer{gin.Default()}

	registerRoutes(server)

	return server
}
