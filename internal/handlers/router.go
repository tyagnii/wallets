package handlers

import (
	"github.com/gin-gonic/gin"
)

// NewRouter creates http server with default options
func NewRouter(h Handler) *gin.Engine {

	r := gin.Default()

	r.POST("/api/v1/wallet", h.PostAmount)
	r.GET("/api/v1/wallets/:UUID", h.GetWalletBalance)
	r.POST("/api/v1/wallets/create", h.CreateWallet)

	return r
}
