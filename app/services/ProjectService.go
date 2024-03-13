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
		projMeta.ProjectId = proj.Id
		projMeta.ProjectName = proj.ProjectName
		projMeta.UserName = proj.UserName
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

func GetProjectDataById(projectId string) models.SingleProjectsResp {
	projectInfo, err := db.GetProjectById(projectId)
	if err != nil {
		resp := models.SingleProjectsResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		return resp
	}
	resp := models.SingleProjectsResp{
		Status:  "SUCCESS",
		Message: "ALL_USER_PROJECTS_FETCHED",
		ProjectData: models.ProjectData{
			ProjectId:   projectInfo.Id,
			ProjectName: projectInfo.ProjectName,
			UserName:    projectInfo.UserName,
			Html:        projectInfo.Html,
			Css:         projectInfo.Css,
			Javascript:  projectInfo.Javascript,
		},
	}
	return resp
}

func UpdateProjectService(projReq models.UpdateProjReq) models.UpdateProjResp {
	err := db.UpdateProject(projReq)
	if err != nil {
		resp := models.UpdateProjResp{
			Status:  "ERROR",
			Message: "DB_ERROR",
		}
		return resp
	}
	return models.UpdateProjResp{
		Status:  "SUCCESS",
		Message: "PROJECT_UPDATED_SUCCESSFULY",
	}
}
