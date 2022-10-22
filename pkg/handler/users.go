package handler

import (
	"net/http"
	"strconv"
	"taskexchange"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	Email     string  `json:"email" binding:"required,max=255"`
	Password  string  `json:"password" binding:"required,max=255"`
	Username  string  `json:"username" binding:"required,max=255"`
	FirstName string  `json:"first_name" binding:"required,max=255"`
	LastName  string  `json:"last_name" binding:"required,max=255"`
	Type      int     `json:"type" binding:"required"`
	Balance   float64 `json:"balance"`
	Points    int     `json:"points"`
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
		Username:   input.Username,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      input.Email,
		Password:   input.Password,
		Type:       input.Type,
		Balance:    input.Balance,
		Points:     input.Points,
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

type getAllUsersResponse struct {
	Data []taskexchange.User `json:"data"`
}

type getHiddenUsersResponse struct {
	Data []taskexchange.UserHidden `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	isAdmin := h.checkIsAdmin(c)

	if isAdmin {
		users, err := h.services.Users.GetAll()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, getAllUsersResponse{
			Data: users,
		})
	} else {
		users, err := h.services.Users.GetAllHidden()
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, getHiddenUsersResponse{
			Data: users,
		})
	}
}

type getOneUserHiddenResponse struct {
	Data taskexchange.UserHidden `json:"data"`
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.Users.GetByIdHidden(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOneUserHiddenResponse{
		Data: user,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	currentUser, err := getUser(c)
	if err != nil {
		return
	}

	if currentUser.Type != 3 {
		newErrorResponse(c, http.StatusBadRequest, "wrong user type")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.Users.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input taskexchange.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Users.Update(user, input)
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
