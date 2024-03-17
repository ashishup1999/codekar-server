package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfilesByNamehandler(c *gin.Context) {
	var req models.ProfilesByNameReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.GetProfilesByNameService(req.Name, req.PageNo))
}

func GetProjectsByNamehandler(c *gin.Context) {
	var req models.ProjectsByNameReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.GetProjectsByName(req.ProjName, req.PageNo))
}

func GetPgsByNamehandler(c *gin.Context) {
	var req models.PgsByNameReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.GetPlaygroundsByName(req.PgName, req.PageNo))
}

func GetWbsByNamehandler(c *gin.Context) {
	var req models.WbsByNameReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.GetWbsByName(req.WbName, req.PageNo))
}
