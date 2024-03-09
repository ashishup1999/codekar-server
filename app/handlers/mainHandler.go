package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUphandler(r *gin.Engine){
	r.GET("/handleRoot",handleRoot)
	r.POST("/compile",CompileCode)
}

func handleRoot(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello from handler!",
    })
}