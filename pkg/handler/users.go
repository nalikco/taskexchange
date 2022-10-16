package handler

import (
	"net/http"
	"strconv"
	"taskexchange"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	Email    string  `json:"email" binding:"required,max=255" db:"email"`
	Password string  `json:"password" binding:"required,max=255"`
	Username string  `json:"username" binding:"required,max=255" db:"username"`
	Type     int     `json:"type" binding:"required" db:"type"`
	Balance  float64 `json:"balance" db:"balance"`
	Points   int     `json:"points" db:"points"`
}

func (h *Handler) createUser(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	if user.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	var input createUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Users.CreateUser(taskexchange.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Type:     input.Type,
		Balance:  input.Balance,
		Points:   input.Points,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllUsersResponse struct {
	Data []taskexchange.User `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Users.GetAll(false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

type getOneUserResponse struct {
	Data             taskexchange.User `json:"data"`
	ActiveTasksCount int               `json:"active_tasks_count"`
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.Users.GetById(id, false)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOneUserResponse{
		Data: user,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
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

	var input taskexchange.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Users.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
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

	err = h.services.Users.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
