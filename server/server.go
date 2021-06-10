package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func (server Server) Run() {
	err := server.engine.Run("127.0.0.1:" + server.port)
	if err != nil {
		log.Fatalln("Cannot start server")
	}
}

func NewServer() *Server {

	var server Server
	server.port = "5000"
	server.engine = gin.Default()

	configRoutes(server.engine)

	return &server
}
