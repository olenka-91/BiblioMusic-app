package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponce struct {
	Message string `json:"message"`
}
type Response struct {
	ID int `json:"id"`
}

type StatusResponce struct {
	Status string `json:"status"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string, errMessage string) {
	logrus.WithField("Err:", errMessage).Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponce{Message: message})
}
