package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetUphandler(r *gin.Engine) {
	r.POST("/compile", CompileCode)
	r.POST("/createUserAccount", CreateUserHandler)
}
