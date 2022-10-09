package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskexchange"
)

type createOptionInput struct {
	ParentId int     `json:"parent_id"`
	Title    string  `json:"title" binding:"required,max=100"`
	Price    float64 `json:"price" binding:"required,numeric"`
}

func (h *Handler) createOption(c *gin.Context) {
	err := checkUserType(c, 3)
	if err != nil {
		return
	}

	var input createOptionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(input.ParentId, taskexchange.Option{
		Title: input.Title,
		Price: input.Price,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllOptionsResponse struct {
	Data []taskexchange.Option `json:"data"`
}

func (h *Handler) getAllOptions(c *gin.Context) {
	options, err := h.services.Option.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOptionsResponse{
		Data: options,
	})
}

type getOneOptionResponse struct {
	Data taskexchange.Option `json:"data"`
}

func (h *Handler) getOptionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	option, err := h.services.Option.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOneOptionResponse{
		Data: option,
	})
}

func (h *Handler) updateOption(c *gin.Context) {
	err := checkUserType(c, 3)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdateOptionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Option.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteOption(c *gin.Context) {
	err := checkUserType(c, 3)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Option.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
