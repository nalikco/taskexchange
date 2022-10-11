package handler

import (
	"net/http"
	"strconv"
	"taskexchange"
	"time"

	"github.com/gin-gonic/gin"
)

type findEventsResponse struct {
	Data []taskexchange.Event `json:"data"`
}

func (h *Handler) pollingEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var after int
	afterParam, ok := c.GetQuery("after")
	if !ok || afterParam == "" {
		after = 0
	} else {
		after, _ = strconv.Atoi(afterParam)
	}

	if after == 0 {
		after, err = h.services.Events.GetLastId(userId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := 0; i < 9; i++ {
		events, err := h.services.Events.Polling(userId, after)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if len(events) != 0 {
			c.JSON(http.StatusOK, findEventsResponse{
				Data: events,
			})
			return
		}

		time.Sleep(time.Second)
	}

	c.JSON(http.StatusOK, findEventsResponse{
		Data: []taskexchange.Event{},
	})
}

func (h *Handler) findNewEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := h.services.Events.GetNew(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findEventsResponse{
		Data: events,
	})
}

type findAllEventsResponse struct {
	Pagination paginationResponse   `json:"pagination"`
	Data       []taskexchange.Event `json:"data"`
}

func (h *Handler) findAllEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	pagination := taskexchange.NewPagination(c, 1, 20)

	events, pagination, err := h.services.Events.GetAll(userId, pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllEventsResponse{
		Pagination: paginationResponse{
			Count:       pagination.Count,
			Pages:       pagination.Pages,
			CurrentPage: pagination.CurrentPage,
			PerPage:     pagination.PerPage,
			Offset:      pagination.Offset,
		},
		Data: events,
	})
}

func (h *Handler) viewAllEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Events.ViewAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) viewEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Events.View(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Events.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
