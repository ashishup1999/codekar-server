package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUserProjshandler(c *gin.Context) {
	userName := c.Param("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetAllProjectsByUser(userName))
}

func CreateNewProject(c *gin.Context) {
	var req models.CreateProjReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.CreateNewProjectByUsername(req))
}

func GetProjectById(c *gin.Context) {
	projId := c.Param("projId")
	if projId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetProjectDataById(projId))
}

func UpdateProjectHandler(c *gin.Context) {
	var req models.UpdateProjReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.UpdateProjectService(req))
}

func GetProjectThumbnailHandler(c *gin.Context) {
	projId := c.Param("projId")
	if projId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	imgBuff, err := services.GetProjectThumbnailService(projId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	base64Encoded := base64.StdEncoding.EncodeToString(imgBuff)
	c.JSON(http.StatusAccepted, gin.H{"status": "SUCCESS", "base64Img": base64Encoded})
}
