package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/v1/users", server.createUser)
	router.POST("/v1/users/login", server.loginUser)
	router.POST("/v1/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/v1").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)
	server.router = router
}
