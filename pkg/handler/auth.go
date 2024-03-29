package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskexchange"
	"time"
)

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type SignUpInput struct {
	Email     string `json:"email" binding:"required,email,min=4,max=70"`
	Password  string `json:"password" binding:"required,min=6,max=100"`
	Username  string `json:"username" binding:"required,min=4,max=40"`
	FirstName string `json:"first_name" binding:"required,min=4,max=40"`
	LastName  string `json:"last_name" binding:"required,min=4,max=40"`
	Type      int    `json:"type" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input SignUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(taskexchange.User{
		Username:   input.Username,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      input.Email,
		Password:   input.Password,
		Type:       input.Type,
		Balance:    0,
		Points:     0,
		LastOnline: time.Now(),
		CreatedAt:  time.Now(),
		DeletedAt:  nil,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getOneUserResponse struct {
	Data             taskexchange.User `json:"data"`
	ActiveTasksCount int               `json:"active_tasks_count"`
}

func (h *Handler) getMyUser(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		return
	}

	activeTasksCount, err := h.services.Tasks.CountActiveByUser(user.Id)

	c.JSON(http.StatusOK, getOneUserResponse{
		Data:             user,
		ActiveTasksCount: activeTasksCount,
	})
}
