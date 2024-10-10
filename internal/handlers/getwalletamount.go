package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWalletAmount
//
// TODO:
// Add db connection
// Hide models at internal/models
// log through gin server logger
func (h *Handler) GetWalletBalance(c *gin.Context) {
	uuid := c.Params.ByName("UUID")
	c.JSON(http.StatusOK, gin.H{"uuid": uuid, "value": "some amount returned"})
}
