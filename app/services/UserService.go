package services

import (
	"codekar/app/db"
	"codekar/app/models"
)

func GetProfilesByNameService(name string, pageNo int64) models.ProfilesResp {
	var resp models.ProfilesResp
	profiles, err := db.GetProfilesByName(name, pageNo)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Profiles = profiles
	resp.Status = "SUCCESS"
	resp.Message = "PROFILES_FETCHED"
	return resp
}

func AddConnection(userName1 string, userName2 string) models.StatusResp {
	var resp models.StatusResp
	err := db.AddUserConnections(userName1, userName2)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "CONNECTION_ADDED_SUCCESSFULY"
	return resp
}

func GetConnectionsByUser(userName string) models.ConnectionsResponse {
	var resp models.ConnectionsResponse
	conns, err := db.GetConnectionsByUser(userName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "CONNECTION_FETCHED_SUCCESSFULY"
	resp.Connections = conns
	return resp
}

func GetUserInfo(userName string) models.UserMetaResp {
	resp, err := db.GetUserInfo(userName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "USER_INFO_FETCHED_SUCCESSFULY"
	return resp
}
