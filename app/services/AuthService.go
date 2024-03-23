package services

import (
	"codekar/app/db"
	"codekar/app/models"
	"codekar/app/utils"
	"fmt"
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
		UserName:    req.UserName,
		Email:       req.Email,
		FullName:    req.FullName,
		Password:    req.Password,
		Connections: []string{},
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
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

func GetVerificationCode(email string) models.AuthenticateUserResp {
	var resp models.AuthenticateUserResp

	//check if user exists
	userExists, err := db.UserExistsByEmail(email)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	} else if !userExists {
		resp.Status = "ERROR"
		resp.Message = "USER_DOES_NOT_EXIST"
		return resp
	}
	//generate otp
	otp, err := db.SetOtp(email)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "OTP_GENERATION_FAILED"
		return resp
	}

	err = utils.SendOtpOnMail(email, otp)
	if err != nil {
		fmt.Println(err.Error())
		resp.Status = "ERROR"
		resp.Message = "UNABLE_TO_SEND_MAIL"
		return resp
	}

	resp.Status = "SUCCESS"
	resp.Message = "VERIFICATION_CODE_GENERATED"
	return resp
}

func ValidateOtpService(req models.OtpReq) models.OtpRes {
	var resp models.OtpRes

	//check if user exists
	err := db.ValidateOtp(req)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = err.Error()
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "OTP_VALIDATION_SUCCESSFUL"
	return resp
}

func UpdatePasswordService(req models.PasswordUpdateReq) models.AuthenticateUserResp {
	var resp models.AuthenticateUserResp
	err := db.UpdateUserPasword(req.Email, req.NewPass)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "PASSWORD_COULD_NOT_BE_UPDATED"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "PASSWORD_UPDATED_SUCCESSFULY"
	return resp
}
