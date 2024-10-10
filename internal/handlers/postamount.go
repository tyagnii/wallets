package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostAmount(c *gin.Context) {
	var pr PostAmountRequest
	if err := c.ShouldBindJSON(&pr); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.Status(http.StatusOK)
}
