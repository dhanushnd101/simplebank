package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/db/sqlc"
)

// server HTTP requests for our bank service
type Server struct{
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP service and the setup router 
func NewServer(store db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
	}

	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)

	router.POST("/transfers",server.createTransfer)

	router.POST("/users",server.createUser)

	// add router to router
	server.router = router
	return server  
}

// start the http server on a specific address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}

