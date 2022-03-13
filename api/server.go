package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/haibuiminh/Simplebank/db/sqlc"
)

// Server serves HTTP request for out banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccount)
	router.DELETE("/api/accounts/:id", server.deleteAccount)
	router.PUT("/api/accounts/:id", server.updateAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
