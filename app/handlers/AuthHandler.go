package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var req models.CreateUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.CreateUserAcc(req))
}

func AuthenticateUserHandler(c *gin.Context) {
	var req models.AuthenticateUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.AuthenticateUser(req))
}

func GetVerificationCodeHandler(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": "Bad Requests"})
		return
	}
	c.JSON(http.StatusAccepted, services.GetVerificationCode(email))
}

func ValidateOtpHandler(c *gin.Context) {
	var req models.OtpReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "ERROR", "message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, services.ValidateOtpService(req))
}
