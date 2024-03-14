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
	r.POST("/createUserAccount", CreateUserHandler)
	r.POST("/authenticateUser", AuthenticateUserHandler)
	r.GET("/getAllProjectsByUser/:userName", GetAllUserProjshandler)
	r.POST("/createNewProjectByUser", CreateNewProject)
	r.GET("/getProjectById/:projId", GetProjectById)
	r.POST("/updateProject", UpdateProjectHandler)
	r.GET("/getProjectThumbnail/:projId", GetProjectThumbnailHandler)
}
