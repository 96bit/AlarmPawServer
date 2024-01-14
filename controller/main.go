package controller

import (
	"AlarmPawServer/push"
	"github.com/gin-gonic/gin"
)

func BaseController(c *gin.Context) {

	params, err := ParamsHandler(c)
	if err != nil {
		return
	}

	err = push.Push(params)

	if err != nil {
		c.JSON(200, failed(500, "push failed: %v", err))
	}

	c.JSON(200, success())

}

func GetInfo(c *gin.Context) {

}
