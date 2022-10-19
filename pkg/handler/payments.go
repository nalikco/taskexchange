package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskexchange"
)

type paymentsResponse struct {
	Data []taskexchange.Payment `json:"data"`
}

func (h *Handler) GetUserPayments(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	payments, err := h.services.Payments.GetByUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, paymentsResponse{
		Data: payments,
	})
}
