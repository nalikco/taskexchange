package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"taskexchange"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
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

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.services.Authorization.UpdateOnline(userId); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "failed while updating online")
		return
	}

	user, err := h.services.Users.GetById(userId, true)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "failed while getting user")
		return
	}

	c.Set(userCtx, user)
}

func getUser(c *gin.Context) (taskexchange.User, error) {
	userFromCtx, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return taskexchange.User{}, errors.New("user id not found")
	}

	user, ok := userFromCtx.(taskexchange.User)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return taskexchange.User{}, errors.New("user id is of invalid type")
	}

	return user, nil
}

func (h *Handler) checkIsAdmin(c *gin.Context) bool {
	header := c.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")
	if len(headerParts) == 2 {
		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err == nil {
			user, err := h.services.Users.GetById(userId, true)
			if err == nil {
				if user.Type == 3 {
					return true
				}
			}
		}
	}

	return false
}
