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
	Short    string  `json:"short" binding:"max=100"`
	Price    float64 `json:"price" binding:"required,numeric"`
}

func (h *Handler) createOption(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	var input createOptionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Options.Create(input.ParentId, taskexchange.Option{
		Title: input.Title,
		Short: &input.Short,
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
	isAdmin := h.checkIsAdmin(c)

	var err error
	var options []taskexchange.Option
	if isAdmin {
		options, err = h.services.Options.GetAll(true)
	} else {
		options, err = h.services.Options.GetAll(false)
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllOptionsResponse{
		Data: options,
	})
}

func (h *Handler) getCategories(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	options, err := h.services.Options.GetCategories()
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

	option, err := h.services.Options.GetById(id, false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOneOptionResponse{
		Data: option,
	})
}

func (h *Handler) updateOption(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
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

	err = h.services.Options.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteOption(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Options.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
