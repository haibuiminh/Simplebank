package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/haibuiminh/Simplebank/db/sqlc"
	"github.com/haibuiminh/Simplebank/token"
	"github.com/haibuiminh/Simplebank/util"
)

// Server serves HTTP request for out banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/api/users", server.createUser)
	router.POST("/api/users/login", server.loginUser)

	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.GET("/api/accounts", server.listAccount)
	router.DELETE("/api/accounts/:id", server.deleteAccount)
	router.PUT("/api/accounts/:id", server.updateAccount)

	router.POST("/api/transfers", server.createTransfer)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
