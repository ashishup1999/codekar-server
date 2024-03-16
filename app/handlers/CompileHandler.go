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
	var output string
	var err error
	if req.Language == JAVA {
		output, err = services.RunJavafn(req)
	} else if req.Language == PYTHON {
		output, err = services.RunPythonfn(req)
	} else if req.Language == CPP {
		output, err = services.RunCppfn(req)
	} else if req.Language == GO {
		output, err = services.RunGofn(req)
	} else if req.Language == JS {
		output, err = services.RunJSfn(req)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Language not supported"})
		return
	}
	if err != nil {
		c.JSON(500, models.CompileResp{Status: "ERROR", Message: "Execution Failed", Output: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, models.CompileResp{Status: "SUCCESS", Message: "Compiled Successfuly", Output: output})
}
