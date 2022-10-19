package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"taskexchange"
	"time"
)

type inputCreateTasks struct {
	Tasks []inputCreateOneTask `json:"tasks" binding:"required"`
}

type inputCreateOneTask struct {
	Status       int    `json:"status" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	DeliveryDate string `json:"delivery_date" binding:"required"`
	Link         string `json:"link" binding:"required,max=255,url"`
	Description  string `json:"description" binding:"required"`
	Options      []int  `json:"options" binding:"required"`
}

func (h *Handler) createTask(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 2 && user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	var tasksInput inputCreateTasks
	if err := c.BindJSON(&tasksInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var tasks []taskexchange.Task

	if len(tasksInput.Tasks) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "wrong tasks length")
		return
	}

	for _, input := range tasksInput.Tasks {
		if len(input.Options) == 0 {
			newErrorResponse(c, http.StatusBadRequest, "wrong options length")
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
			option, err := h.services.Options.GetById(optionId, true)
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
			CustomerId:   user.Id,
			Status:       input.Status,
			Amount:       input.Amount,
			DeliveryDate: deliveryDate,
			Link:         input.Link,
			Description:  input.Description,
			Options:      options,
		}

		tasks = append(tasks, task)
	}

	amountTasksPrice := 0.00
	for _, task := range tasks {
		_, err = h.services.Tasks.Create(task)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		amountTasksPrice += task.CalculatePrice()
	}

	payment := taskexchange.Payment{
		User: user,
		Type: 1,
		Comment: "Создание задач",
		Sum: amountTasksPrice,
	}
	_, err = h.services.Payments.Create(payment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) createTaskFromExcelFile(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 2 && user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	extension := filepath.Ext(file.Filename)
	if extension != ".xlsx" {
		newErrorResponse(c, http.StatusBadRequest, "wrong file extension")
		return
	}

	filename := fmt.Sprintf("uploads/tmp/%d-%s", time.Now().Unix(), file.Filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	amountTasksPrice, err := h.services.Tasks.CreateFromExcelFile(user.Id, filename)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payment := taskexchange.Payment{
		User: user,
		Type: 1,
		Comment: "Создание задач",
		Sum: amountTasksPrice,
	}
	_, err = h.services.Payments.Create(payment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type getAllTasksResponse struct {
	Pagination paginationResponse  `json:"pagination"`
	Data       []taskexchange.Task `json:"data"`
}

func (h *Handler) getAllTasks(c *gin.Context) {
	isAdmin := h.checkIsAdmin(c)
	userId := 0

	pagination := taskexchange.NewPagination(c, 1, 20)

	if isAdmin {
		userId = -1
	}
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
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 2 && user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "invalid user type")
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

	if input.Options != nil && len(*input.Options) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "wrong options length")
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
	if task.DeletedAt != nil {
		newErrorResponse(c, http.StatusInternalServerError, "task deleted")
		return
	}
	if task.CustomerId != user.Id {
		newErrorResponse(c, http.StatusBadRequest, "wrong user id")
		return
	}

	if task.Status != 0 {
		if input.Status != nil {
			_, err = h.services.Tasks.Update(id, taskexchange.UpdateTaskInput{
				Status: input.Status,
			})
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})

		return
	}

	amountTasksPrice, err := h.services.Tasks.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if amountTasksPrice != 0 {
		var payment taskexchange.Payment
		if amountTasksPrice < 0 {
			payment = taskexchange.Payment{
				User: user,
				Type: 2,
				Comment: "Редактирование задач",
				Sum: math.Abs(amountTasksPrice),
			}
		} else {
			payment = taskexchange.Payment{
				User: user,
				Type: 1,
				Comment: "Редактирование задач",
				Sum: amountTasksPrice,
			}
		}

		_, err = h.services.Payments.Create(payment)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteTask(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 2 && user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
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
	if task.CustomerId != user.Id && user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user id")
		return
	}
	task.Customer, err = h.services.Users.GetById(task.CustomerId, true)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Tasks.Delete(id, task, user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payment := taskexchange.Payment{
		User: user,
		Type: 2,
		Comment: "Удаление задач",
		Sum: task.CalculatePrice(),
	}
	_, err = h.services.Payments.Create(payment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
