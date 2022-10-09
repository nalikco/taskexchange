package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	userTypeCtx         = "userType"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, userType, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.services.Authorization.UpdateOnline(userId); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "failed while updating online")
		return
	}

	c.Set(userIdCtx, userId)
	c.Set(userTypeCtx, userType)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getUserType(c *gin.Context) (int, error) {
	userType, ok := c.Get(userTypeCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user type not found")
		return 0, errors.New("user type not found")
	}

	userTypeInt, ok := userType.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user type is of invalid type")
		return 0, errors.New("user type is of invalid type")
	}

	return userTypeInt, nil
}

func checkUserType(c *gin.Context, neededType int) error {
	userType, err := getUserType(c)
	if err != nil {
		return errors.New("user type is invalid")
	}

	if userType != neededType {
		newErrorResponse(c, http.StatusInternalServerError, "access denied")
		return errors.New("access denied")
	}

	return nil
}
