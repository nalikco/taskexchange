package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskexchange"
	"time"
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
		after, err = h.services.Events.GetLastUserEventId(userId)
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
		events, err := h.services.Events.PollingEvents(userId, after)
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
	return
}

func (h *Handler) findNewEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	events, err := h.services.Events.GetNewEvents(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findEventsResponse{
		Data: events,
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

	err = h.services.Events.ViewEvent(userId, id)
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

	err = h.services.Events.DeleteEvent(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
