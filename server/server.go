package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/z0mi3ie/goimgs/config"
	"github.com/z0mi3ie/goimgs/db"
	"github.com/z0mi3ie/goimgs/routers"
)

// Server is the instance of the server running and also holds the server wide
// configuration and fields
type Server struct {
	router *gin.Engine
}

// NewServer creates a new server struct and performs initial setup
func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}
	return server
}

// Start starts the server
func (s *Server) Start() {
	dbClient, err := db.NewMySQLClient(config.MySQLUser, config.MySQLPassword, config.MySQLDatabase)
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	defer dbClient.DB().Close()

	s.router.MaxMultipartMemory = config.ServerMaxFileSize

	s.router.POST("/image", routers.UploadImage)
	s.router.GET("/images", routers.GetImages)

	s.router.Run(fmt.Sprintf(":%v", config.ServerPort)) // listen and serve on 0.0.0.0:8080
}
