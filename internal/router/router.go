package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/wallets/internal/handlers"
)

// NewRouter creates http server with default options
func NewRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/api/v1/wallet", handlers.PostAmount)
	r.GET("/api/v1/wallets/:UUID", handlers.GetWalletBalance)

	return r
}
