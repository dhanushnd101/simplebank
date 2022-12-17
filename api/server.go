package api

import (
	"github.com/techschool/simplebank/util"
	"fmt"
	
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
)

// server HTTP requests for our bank service
type Server struct{
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP service and the setup router 
func NewServer(config util.Config, store db.Store) (*Server, error){
	// We have two implementations for encrypting 
	// 1. NewPasetoMaker
	// 2. NewJWTMaker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil{
		return nil, fmt.Errorf("Cannot create token maker:%w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}
	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
	}

	server.setupRouter()
	return server, nil 
}

func (server *Server) setupRouter(){

	router := gin.Default()

	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)

	router.POST("/transfers",server.createTransfer)

	router.POST("/users",server.createUser)
	router.POST("/users/login",server.loginUser)

	// add router to router
	server.router = router
}
// start the http server on a specific address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}

