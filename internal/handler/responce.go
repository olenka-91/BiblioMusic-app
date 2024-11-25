package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string, errMessage string) {
	logrus.WithField("Err:", errMessage).Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponce{Message: message})
}
