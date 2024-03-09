package handlers

import (
	"codekar/app/models"
	"codekar/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	JAVA   = "java"
	PYTHON = "python"
	CPP    = "cpp"
	GO     = "go"
	JS     = "javascript"
)

func CompileCode(c *gin.Context) {
	var req models.CompileReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var resp models.CompileResp
	if req.Language == JAVA {
		resp = services.RunJavafn(req)
	} else if req.Language == PYTHON {
		resp = services.RunPythonfn(req)
	} else if req.Language == CPP {
		resp = services.RunCppfn(req)
	} else if req.Language == GO {
		resp = services.RunGofn(req)
	} else if req.Language == JS {
		resp = services.RunJSfn(req)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not supported"})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
