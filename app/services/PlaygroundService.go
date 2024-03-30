package services

import (
	"codekar/app/db"
	"codekar/app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllPlaygroundsByUser(userName string) models.AllPgsResp {
	var resp models.AllPgsResp
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
	pgs, err := db.GetAllPgs(userName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	for _, pg := range pgs {
		pgData := models.PgMeta{
			PgId:     pg.Id,
			PgName:   pg.PgName,
			UserName: pg.UserName,
		}
		resp.Pgs = append(resp.Pgs, pgData)
	}
	resp.Status = "SUCCESS"
	resp.Message = "ALL_USER_PGS_FETCHED"
	return resp
}

func CreateNewPgByUsername(req models.CreatePgReq) models.CreatePgResp {
	var resp models.CreatePgResp
	id, err := db.CreateNewPg(req)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "PG_CREATED_SUCCESSFULY"
	resp.PgId = id
	return resp
}

func GetPgDataById(pgId string) models.SinglePgResp {
	pgInfo, err := db.GetPgById(pgId)
	if err != nil {
		resp := models.SinglePgResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		if err.Error() == mongo.ErrNoDocuments.Error() {
			resp.Message = "PG_DOES_NOT_EXISTS"
		}
		return resp
	}
	resp := models.SinglePgResp{
		Status:  "SUCCESS",
		Message: "ALL_USER_PGS_FETCHED",
		PgData:  pgInfo,
	}
	return resp
}

func UpdatePgService(pgReq models.UpdatePgReq) models.UpdatePgResp {
	err := db.UpdatePg(pgReq)
	if err != nil {
		resp := models.UpdatePgResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		return resp
	}
	return models.UpdatePgResp{
		Status:  "SUCCESS",
		Message: "PG_UPDATED_SUCCESSFULY",
	}
}

func GetPlaygroundsByName(pgName string, pageNo int64) models.AllPgsResp {
	var resp models.AllPgsResp
	pgs, err := db.GetPgsByName(pgName, pageNo)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	for _, pg := range pgs {
		pgData := models.PgMeta{
			PgId:     pg.Id,
			PgName:   pg.PgName,
			UserName: pg.UserName,
		}
		resp.Pgs = append(resp.Pgs, pgData)
	}
	resp.Status = "SUCCESS"
	resp.Message = "PGS_FETCHED"
	return resp
}

func DeletePgService(pgId string) models.UpdateProjResp {
	err := db.DeletePg(pgId)
	if err != nil {
		resp := models.UpdateProjResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		return resp
	}
	return models.UpdateProjResp{
		Status:  "SUCCESS",
		Message: "PG_DELETED_SUCCESSFULY",
	}
}
