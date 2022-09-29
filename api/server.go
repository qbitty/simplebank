package api

import (
	"github.com/gin-gonic/gin"
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
	router.POST("/v1/accounts", server.CreateAccount)
	router.GET("/v1/accounts/:id", server.GetAccount)
	router.GET("/v1/accounts", server.ListAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
