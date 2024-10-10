package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tyagnii/wallets/config"
)

func (h *Handler) CreateWallet(c *gin.Context) {
	if id, err := h.dbConnector.CreateWallet(c); err != nil {
		c.String(http.StatusBadRequest, err.Error()+config.ConnectionString)
	} else {
		c.JSON(http.StatusOK, id)
	}
}
