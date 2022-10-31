package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/v1/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/v1/accounts", server.createAccount)
	router.GET("/v1/accounts/:id", server.getAccount)
	router.GET("/v1/accounts", server.listAccount)

	router.POST("/v1/transfers", server.createTransfer)
	server.router = router
}
