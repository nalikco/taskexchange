package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskexchange"
)

type findOffersResponse struct {
	Data []taskexchange.Offer `json:"data"`
}

func (h *Handler) GetPerformerActiveOffers(c *gin.Context) {
	err := checkUserType(c, 1)
	if err != nil {
		return
	}
	performerId, err := getUserId(c)
	if err != nil {
		return
	}

	offers, err := h.services.Offers.GetPerformerActive(performerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findOffersResponse{
		Data: offers,
	})

}

func (h *Handler) CreateOffer(c *gin.Context) {
	err := checkUserType(c, 1)
	if err != nil {
		return
	}
	performerId, err := getUserId(c)
	if err != nil {
		return
	}

	var input taskexchange.CreateOfferInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offerId, err := h.services.Offers.Make(performerId, input.TaskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": offerId,
	})
}

func (h *Handler) UpdateOffer(c *gin.Context) {
	err := checkUserType(c, 2)
	if err != nil {
		return
	}
	customerId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdateOfferInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Status == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "wrong status")
		return
	}

	err = h.services.Offers.ChangeStatus(id, customerId, input.Status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
