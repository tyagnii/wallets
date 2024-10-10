package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/wallets/internal/handlers"
)

// NewRouter creates http server with default options
func NewRouter() *gin.Engine {

	// add error handling
	h, _ := handlers.NewHandler()

	r := gin.Default()

	r.POST("/api/v1/wallet", h.PostAmount)
	r.GET("/api/v1/wallets/:UUID", h.GetWalletBalance)

	return r
}
