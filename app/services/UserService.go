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
