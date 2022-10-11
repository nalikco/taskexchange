package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

type paginationResponse struct {
	Pages       int `json:"pages"`
	Count       int `json:"count"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Offset      int `json:"offset"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
