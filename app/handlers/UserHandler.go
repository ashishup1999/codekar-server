package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConnectionReqHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.ConnectionReqService(req.Sender, req.Reciever))
}

func ConnectionStatusHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.ConnectionStatusService(req.Sender, req.Reciever))
}

func GetAllConnectionReqsHandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": "BAD_REQUEST"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetAllConnectionReqsService(userName))
}

func RejectConnectionReqHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.RejectConnectionReqService(req.Reciever, req.Sender))
}

func AddConnectionHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.AddConnection(req.Reciever, req.Sender))
}

func RemoveConnectionHandler(c *gin.Context) {
	var req models.AddConnectionReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.RemoveConnection(req.Reciever, req.Sender))
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

func DeleteWbHandler(c *gin.Context) {
	wbId := c.Param("wbId")
	if wbId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.DeleteWbService(wbId))
}
