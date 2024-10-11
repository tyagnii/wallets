package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetWalletBalance returns wallet balance
func (h *Handler) GetWalletBalance(c *gin.Context) {
	var wb WalelteBalance

	wb.WalletID = c.Params.ByName("UUID")

	b, err := h.dbConnector.GetWalletBalance(c, wb.WalletID)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	wb.Balance = strconv.Itoa(b)
	// wbJSON, err := json.Marshal(wb)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, wb)
}
