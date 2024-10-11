package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWalletBalance returns wallet balance
func (h *Handler) GetWalletBalance(c *gin.Context) {
	uuid := c.Params.ByName("UUID")
	c.JSON(http.StatusOK, gin.H{"uuid": uuid, "value": "some amount returned"})
}
