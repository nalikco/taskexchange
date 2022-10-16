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
	user, err := getUser(c)
	if err != nil {
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
		after, err = h.services.Events.GetLastId(user.Id)
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
		events, err := h.services.Events.Polling(user.Id, after)
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
	user, err := getUser(c)
	if err != nil {
		return
	}

	events, err := h.services.Events.GetNew(user.Id)
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
	user, err := getUser(c)
	if err != nil {
		return
	}

	pagination := taskexchange.NewPagination(c, 1, 20)

	events, pagination, err := h.services.Events.GetAll(user.Id, pagination)
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
	user, err := getUser(c)
	if err != nil {
		return
	}

	err = h.services.Events.ViewAll(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) viewEvent(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Events.View(user.Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteEvent(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Events.Delete(user.Id, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
