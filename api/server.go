package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "pro.qbitty/simplebank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(validator.Validate); ok {
		v.RegisterValidation("currency", currencyValid)
	}

	router.POST("/v1/accounts", server.createAccount)
	router.GET("/v1/accounts/:id", server.getAccount)
	router.GET("/v1/accounts", server.listAccount)

	router.POST("/v1/transfers", server.createTransfer)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
