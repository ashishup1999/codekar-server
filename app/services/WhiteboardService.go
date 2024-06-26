package services

import (
	"codekar/app/db"
	"codekar/app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllWbsByUser(userName string) models.AllWbsResp {
	var resp models.AllWbsResp
	validUsername, err := db.UserExistsByUsername(userName)
	if !validUsername {
		resp.Status = "ERROR"
		resp.Message = "USER_DOES_NOT_EXISTS"
		return resp
	} else if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	wbs, err := db.GetAllWbs(userName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	for _, wb := range wbs {
		wbData := models.WbMeta{
			WbId:     wb.Id,
			WbName:   wb.WbName,
			UserName: wb.UserName,
		}
		resp.Wbs = append(resp.Wbs, wbData)
	}
	resp.Status = "SUCCESS"
	resp.Message = "ALL_USER_WBS_FETCHED"
	return resp
}

func CreateNewWbByUsername(req models.CreateWbReq) models.CreateWbResp {
	var resp models.CreateWbResp
	id, err := db.CreateNewWb(req)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "WB_CREATED_SUCCESSFULY"
	resp.WbId = id
	return resp
}

func GetWbDataById(wbId string) models.SingleWbResp {
	wbInfo, err := db.GetWbById(wbId)
	if err != nil {
		resp := models.SingleWbResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		if err.Error() == mongo.ErrNoDocuments.Error() {
			resp.Message = "WB_DOES_NOT_EXISTS"
		}
		return resp
	}
	resp := models.SingleWbResp{
		Status:  "SUCCESS",
		Message: "ALL_USER_WBS_FETCHED",
		WbData:  wbInfo,
	}
	return resp
}

func UpdateWbService(wbReq models.UpdateWbReq) models.UpdateWbResp {
	err := db.UpdateWb(wbReq)
	if err != nil {
		resp := models.UpdateWbResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		return resp
	}
	return models.UpdateWbResp{
		Status:  "SUCCESS",
		Message: "WB_UPDATED_SUCCESSFULY",
	}
}

func GetWbsByName(wbName string, pageNo int64) models.AllWbsResp {
	var resp models.AllWbsResp
	wbs, err := db.GetWbsByName(wbName, pageNo)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	for _, wb := range wbs {
		wbData := models.WbMeta{
			WbId:     wb.Id,
			WbName:   wb.WbName,
			UserName: wb.UserName,
		}
		resp.Wbs = append(resp.Wbs, wbData)
	}
	resp.Status = "SUCCESS"
	resp.Message = "WBS_FETCHED"
	return resp
}
