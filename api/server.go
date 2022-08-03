package api

import (
	db "github.com/ebaudet/go-fizz-buzz/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	server.setupRouter()

	return server
}

// Define all the api's routes here
func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/fizzbuzz", server.fizzBuzz)
	router.GET("/statistics", server.statistics)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
