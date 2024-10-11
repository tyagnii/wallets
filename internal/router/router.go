package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tyagnii/wallets/internal/handlers"
)

// NewRouter creates http server with default options
func NewRouter() *gin.Engine {

	// add error handling
	h, err := handlers.NewHandler()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/api/v1/wallet", h.PostAmount)
	r.GET("/api/v1/wallets/:UUID", h.GetWalletBalance)
	r.POST("/api/v1/wallets/create", h.CreateWallet)

	return r
}
