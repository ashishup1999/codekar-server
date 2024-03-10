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
		resp.Message = "DB_ERROR"
		return resp
	} else if userExists {
		resp.Status = "ERROR"
		resp.Message = "USER_ALREADY_EXISTS"
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
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "USER_CREATED_SUCCESSFULY"
	return resp
}

func AuthenticateUser(req models.AuthenticateUserReq) models.AuthenticateUserResp {
	var resp models.AuthenticateUserResp
	//check if user exists
	userExists, err := db.UserExistsByUsername(req.UserName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	} else if !userExists {
		resp.Status = "ERROR"
		resp.Message = "USER_DOES_NOT_EXIST"
		return resp
	}
	//if user exists now validate tis credentials are right
	isValidated, err := db.ValidateUser(req.UserName, req.Password)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	} else if isValidated {
		resp.Status = "SUCCESS"
		resp.Message = "USER_VALIDATED"
		return resp
	}
	resp.Status = "ERROR"
	resp.Message = "FALSE_CREDS"
	return resp
}
