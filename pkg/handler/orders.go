package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskexchange"
)

func (h *Handler) getAllUserOrders(c *gin.Context) {

}

type findOrdersResponse struct {
	Data []taskexchange.Order `json:"data"`
}

func (h *Handler) getAllPerformerActiveOrders(c *gin.Context) {
	err := checkUserType(c, 1)
	if err != nil {
		return
	}
	performerId, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.Orders.FindActiveByPerformerId(performerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getOrderById(c *gin.Context) {

}

func (h *Handler) updateOrder(c *gin.Context) {

}
