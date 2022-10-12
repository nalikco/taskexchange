package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskexchange"
	"time"
)

type inputCreateTask struct {
	Status       int    `json:"status" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	DeliveryDate string `json:"delivery_date" binding:"required"`
	Link         string `json:"link" binding:"required,max=255,url"`
	Description  string `json:"description" binding:"required"`
	Options      []int  `json:"options" binding:"required"`
}

func (h *Handler) createTask(c *gin.Context) {
	err := checkUserType(c, 2)
	if err != nil {
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input inputCreateTask
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Status != 0 && input.Status != 1 {
		newErrorResponse(c, http.StatusBadRequest, "wrong status")
		return
	}

	if input.Amount <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "wrong amount")
		return
	}

	deliveryDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", input.DeliveryDate))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var options []taskexchange.Option

	for _, optionId := range input.Options {
		option, err := h.services.Options.GetById(optionId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("wrong option id: %d", optionId))
			return
		}

		if option.ParentId != nil {
			var parentIdFound = false

			for _, parentId := range input.Options {
				if parentId == *option.ParentId {
					parentIdFound = true
				}
			}

			if !parentIdFound {
				newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("parent id %d not found in options array for option: %d", *option.ParentId, optionId))
				return
			}
		}

		options = append(options, option)
	}

	task := taskexchange.Task{
		CustomerId:   userId,
		Status:       input.Status,
		Amount:       input.Amount,
		DeliveryDate: deliveryDate,
		Link:         input.Link,
		Description:  input.Description,
		Options:      options,
	}

	id, err := h.services.Tasks.Create(task)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTasksResponse struct {
	Pagination paginationResponse  `json:"pagination"`
	Data       []taskexchange.Task `json:"data"`
}

func (h *Handler) getAllTasks(c *gin.Context) {
	pagination := taskexchange.NewPagination(c, 1, 20)

	tasks, pagination, err := h.services.Tasks.GetAll(0, pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponse{
		Pagination: paginationResponse{
			Count:       pagination.Count,
			Pages:       pagination.Pages,
			CurrentPage: pagination.CurrentPage,
			PerPage:     pagination.PerPage,
			Offset:      pagination.Offset,
		},
		Data: tasks,
	})
}

func (h *Handler) getUserAllTasks(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	pagination := taskexchange.NewPagination(c, 1, 20)

	tasks, pagination, err := h.services.Tasks.GetAll(userId, pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponse{
		Pagination: paginationResponse{
			Count:       pagination.Count,
			Pages:       pagination.Pages,
			CurrentPage: pagination.CurrentPage,
			PerPage:     pagination.PerPage,
			Offset:      pagination.Offset,
		},
		Data: tasks,
	})
}

type oneTaskResponse struct {
	Data taskexchange.Task `json:"data"`
}

func (h *Handler) getTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	task, err := h.services.Tasks.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, oneTaskResponse{
		Data: task,
	})
}

func (h *Handler) updateTask(c *gin.Context) {
	err := checkUserType(c, 2)
	if err != nil {
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.DeliveryDateString != nil {
		deliveryDate, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", *input.DeliveryDateString))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		input.DeliveryDate = &deliveryDate
	}

	task, err := h.services.Tasks.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if task.CustomerId != userId {
		newErrorResponse(c, http.StatusBadRequest, "wrong user id")
		return
	}

	err = h.services.Tasks.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	err := checkUserType(c, 2)
	if err != nil {
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	task, err := h.services.Tasks.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if task.CustomerId != userId {
		newErrorResponse(c, http.StatusBadRequest, "wrong user id")
		return
	}

	err = h.services.Tasks.Delete(id, task, userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
