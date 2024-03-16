package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUserWbshandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetAllWbsByUser(userName))
}

func CreateNewWb(c *gin.Context) {
	var req models.CreateWbReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.CreateNewWbByUsername(req))
}

func GetWbById(c *gin.Context) {
	wbId := c.Param("wbId")
	if wbId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetWbDataById(wbId))
}

func UpdateWbHandler(c *gin.Context) {
	var req models.UpdateWbReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.UpdateWbService(req))
}
