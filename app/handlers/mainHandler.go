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

	//auth apis
	r.POST("/createUserAccount", CreateUserHandler)
	r.POST("/authenticateUser", AuthenticateUserHandler)
	r.GET("/getVerificationCode/:email", GetVerificationCodeHandler)
	r.POST("/validateOtp", ValidateOtpHandler)
	r.POST("/updatePassword", UpdatePasswordHandler)

	//project apis
	r.GET("/getAllProjectsByUser/:userName", GetAllUserProjshandler)
	r.POST("/createNewProjectByUser", CreateNewProject)
	r.GET("/getProjectById/:projId", GetProjectById)
	r.POST("/updateProject", UpdateProjectHandler)
	r.GET("/deleteProject/:projId", DeleteProjectHandler)

	//pg apis
	r.GET("/getAllPgsByUser/:userName", GetAllUserPgshandler)
	r.POST("/createNewPgByUser", CreateNewPg)
	r.GET("/getPgById/:pgId", GetPgById)
	r.POST("/updatePg", UpdatePgHandler)
	r.GET("/deletePg/:pgId", DeletePgHandler)

	//wb apis
	r.GET("/getAllWbsByUser/:userName", GetAllUserWbshandler)
	r.POST("/createNewWbByUser", CreateNewWb)
	r.GET("/getWbById/:wbId", GetWbById)
	r.POST("/updateWb", UpdateWbHandler)
	r.GET("/deleteWb/:wbId", DeleteWbHandler)

	//explore apis
	r.POST("/getProfilesByName", GetProfilesByNamehandler)
	r.POST("/getProjsByName", GetProjectsByNamehandler)
	r.POST("/getPgsByName", GetPgsByNamehandler)
	r.POST("/getWbsByName", GetWbsByNamehandler)

	//user apis
	r.POST("/updateUserDetails", UpdateUserDetailsHandler)
	r.POST("/connectionReq", ConnectionReqHandler)
	r.GET("/getAllConnectionReqs/:userName", GetAllConnectionReqsHandler)
	r.POST("/connectionStatus", ConnectionStatusHandler)
	r.POST("/rejectConnectionReq", RejectConnectionReqHandler)
	r.POST("/addConnection", AddConnectionHandler)
	r.POST("/removeConnection", RemoveConnectionHandler)
	r.GET("/connectionByUser/:userName", GetConnectionsByUserHandler)
	r.GET("/userInfo/:userName", GetUserInfoHandler)

	//file handlers
	r.POST("/uploadProfileImg",UploadProfileImgHandler)
}
