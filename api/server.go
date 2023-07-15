package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/imad-elbouhati/bank/db/sqlc"
	"github.com/imad-elbouhati/bank/token"
	"github.com/imad-elbouhati/bank/util"
)

type Server struct {
	config util.Config
	tokenMaker token.Maker
	store db.Store
	router *gin.Engine
}


func NewServer(config util.Config ,store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWTKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token %w", err)
	}

	server := &Server{store: store,config: config, tokenMaker: tokenMaker}
	router := gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency",validCurrency)
	}


	authRoutes := router.Group("/").Use(authMiddeleware(server.tokenMaker))

	authRoutes.POST("/accounts",server.createAccount)
	authRoutes.GET("/accounts/:id",server.getAccount)
	authRoutes.GET("/accounts/",server.listAccount)
	authRoutes.POST("/transfers",server.createTransfer)

	router.POST("users",server.createUser)
	router.POST("users/login",server.loginUser)

	server.router = router
	return server, nil
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}