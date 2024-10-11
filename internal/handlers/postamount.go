package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostAmount(c *gin.Context) {
	var pr PostAmountRequest
	if err := c.ShouldBindJSON(&pr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	a, err := strconv.Atoi(pr.Amount)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.dbConnector.ChangeWalletBalance(c, pr.WalletID, pr.OperationType, a)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
