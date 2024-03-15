package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUserPgshandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetAllPlaygroundsByUser(userName))
}

func CreateNewPg(c *gin.Context) {
	var req models.CreatePgReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.CreateNewPgByUsername(req))
}

func GetPgById(c *gin.Context) {
	pgId := c.Param("pgId")
	if pgId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetPgDataById(pgId))
}

func UpdatePgHandler(c *gin.Context) {
	var req models.UpdatePgReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.UpdatePgService(req))
}
