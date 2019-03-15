package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/routers"
)

// Server is the instance of the server running and also holds the server wide
// configuration and fields
type Server struct {
	router *gin.Engine
	config *config.Config
}

// NewServer creates a new server struct and performs initial setup
func NewServer() *Server {
	cfg, err := config.NewServerConfig()
	if err != nil {
		// If we can't even initialize the config we are going to have a bad time
		panic(1)
	}
	server := &Server{
		router: gin.Default(),
		config: cfg,
	}
	server.router.MaxMultipartMemory = cfg.ServerMaxFileSizeInt64()

	return server
}

// Start starts the server
func (s *Server) Start() {
	s.router.POST("/images",
		routers.ConfigMiddleware(s.config),
		routers.MySQLClientMiddleware,
		routers.CORSHeaderMiddleware,
		routers.UploadImages,
	)
	s.router.GET("/images",
		routers.ConfigMiddleware(s.config),
		routers.MySQLClientMiddleware,
		routers.CORSHeaderMiddleware,
		routers.GetImages,
	)
	s.router.DELETE("/images",
		routers.ConfigMiddleware(s.config),
		routers.MySQLClientMiddleware,
		routers.DeleteImageQueryParamsMiddleware,
		routers.CORSHeaderMiddleware,
		routers.DeleteImages,
	)

	s.router.Run(fmt.Sprintf(":%v", s.config.ServerPort))
}
