package services

import (
	"codekar/app/db"
	"codekar/app/models"
	"time"
)

func CreateUserAcc(req models.CreateUserReq) models.CreateUserResp {
	var resp models.CreateUserResp
	//check if user exists
	userExists, err := db.UserExistsByUsernameEmail(req.UserName, req.Email)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB Error"
		return resp
	} else if userExists {
		resp.Status = "ERROR"
		resp.Message = "User Already Exists"
		return resp
	}
	//if not than create user and return success
	userObj := models.User{
		UserName:  req.UserName,
		Email:     req.Email,
		FullName:  req.FullName,
		Password:  req.Password,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	err = db.CreateUser(userObj)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB error on creation"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "User Created Successfully"
	return resp
}

func AuthenticateUser(req models.AuthenticateUserReq) models.AuthenticateUserResp {
	var resp models.AuthenticateUserResp
	//check if user exists
	userExists, err := db.UserExistsByUsername(req.UserName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB Error"
		return resp
	} else if !userExists {
		resp.Status = "ERROR"
		resp.Message = "User Does Not Exists"
		return resp
	}
	//if user exists now validate tis credentials are right
	isValidated, err := db.ValidateUser(req.UserName, req.Password)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB Validate Error"
		return resp
	} else if isValidated {
		resp.Status = "SUCCESS"
		resp.Message = "User Validated Successfully"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "False credentials"
	return resp
}
