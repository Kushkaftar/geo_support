package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorStruct struct {
	Error errorResponse
}
type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, massage string) {

	logrus.Error(massage)
	c.AbortWithStatusJSON(statusCode, errorStruct{errorResponse{massage}})
}
