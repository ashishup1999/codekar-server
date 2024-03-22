package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddConnectionHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.AddConnection(req.UserName1, req.UserName2))
}

func GetConnectionsByUserHandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": "BAD_REQUEST"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetConnectionsByUser(userName))
}

func GetUserInfoHandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": "BAD_REQUEST"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetUserInfo(userName))
}
