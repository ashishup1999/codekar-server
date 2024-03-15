package handlers

import (
	"codekar/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetUphandler(r *gin.Engine) {
	//assign middlewares to the routes
	r.Use(middleware.CORSMiddleware())

	//route handling
	r.POST("/compile", CompileCode)

	//user apis
	r.POST("/createUserAccount", CreateUserHandler)
	r.POST("/authenticateUser", AuthenticateUserHandler)

	//project apis
	r.GET("/getAllProjectsByUser/:userName", GetAllUserProjshandler)
	r.POST("/createNewProjectByUser", CreateNewProject)
	r.GET("/getProjectById/:projId", GetProjectById)
	r.POST("/updateProject", UpdateProjectHandler)

	//pg apis
	r.GET("/getAllPgsByUser/:userName", GetAllUserPgshandler)
	r.POST("/createNewPgByUser", CreateNewPg)
	r.GET("/getPgById/:pgId", GetPgById)
	r.POST("/updatePg", UpdatePgHandler)
}
