package services

import (
	"codekar/app/db"
	"codekar/app/models"
)

func GetAllProjectsByUser(userName string) models.AllProjectsResp {
	var resp models.AllProjectsResp
	projects, err := db.GetAllProjects(userName)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	for _, proj := range projects {
		var projMeta models.ProjectMeta
		projMeta.ProjectId = proj.Id.String()
		projMeta.ProjectName = proj.ProjectName
		projMeta.UserName = proj.UserName
		projMeta.Likes = proj.Likes
		projMeta.Comments = proj.Comments
		resp.Projects = append(resp.Projects, projMeta)
	}
	resp.Status = "SUCCESS"
	resp.Message = "ALL_USER_PROJECTS_FETCHED"
	return resp
}

func CreateNewProjectByUsername(req models.CreateProjReq) models.CreateProjResp {
	var resp models.CreateProjResp
	id, err := db.CreateNewProj(req)
	if err != nil {
		resp.Status = "ERROR"
		resp.Message = "DB_ERROR"
		return resp
	}
	resp.Status = "SUCCESS"
	resp.Message = "PROJECT_CREATED_SUCCESSFULY"
	resp.ProjectId = id
	return resp
}
