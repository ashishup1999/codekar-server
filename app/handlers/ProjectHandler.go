package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
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

func DeleteProjectHandler(c *gin.Context) {
	projId := c.Param("projId")
	if projId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR"})
		return
	}
	c.JSON(http.StatusAccepted, services.DeleteProjectService(projId))
}
