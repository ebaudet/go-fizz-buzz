package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests for our service
type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}

	server.setupRouter()

	return server
}

// Define all the api's routes here
func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/fizzbuzz", server.fizzBuzz)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
