package handler

import (
	"net/http"
	"strconv"
	"taskexchange"

	"github.com/gin-gonic/gin"
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

func (h *Handler) getAllPerformerOrders(c *gin.Context) {
	err := checkUserType(c, 1)
	if err != nil {
		return
	}
	performerId, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.Orders.FindAllByPerformerId(performerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findOrdersResponse{
		Data: orders,
	})
}

func (h *Handler) getAllCustomerOrders(c *gin.Context) {
	err := checkUserType(c, 2)
	if err != nil {
		return
	}
	customerId, err := getUserId(c)
	if err != nil {
		return
	}

	orders, err := h.services.Orders.FindAllByCustomerId(customerId)
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
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdateOrderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Orders.Update(id, userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
