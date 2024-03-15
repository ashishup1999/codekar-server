package models

type Project struct {
	Id          string `bson:"_id,omitempty"`
	UserName    string `bson:"username"`
	ProjectName string `bson:"projectname"`
	Html        string `bson:"html"`
	Javascript  string `bson:"javascript"`
	Css         string `bson:"css"`  
	CreatedAt   string `bson:"createdat"`
	UpdatedAt   string `bson:"updatedat"`
}

type ProjectMeta struct {
	ProjectId   string `json:"projectId"`
	UserName    string `json:"userName"`
	ProjectName string `json:"projectName"`
	PreviewHtml string `json:"previewHtml"`
}

type ProjectData struct {
	ProjectId   string `json:"projectId"`
	UserName    string `json:"userName"`
	Html        string `json:"html"`
	Javascript  string `json:"javascript"`
	Css         string `json:"css"`
	ProjectName string `json:"projectName"`
}

type AllProjectsResp struct {
	Status   string        `json:"status"`
	Message  string        `json:"message"`
	Projects []ProjectMeta `json:"projects"`
}

type SingleProjectsResp struct {
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	ProjectData ProjectData `json:"projectData"`
}

type CreateProjReq struct {
	UserName    string `json:"userName"`
	ProjectName string `json:"projectName"`
}

type CreateProjResp struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ProjectId string `json:"projectId"`
}

type UpdateProjReq struct {
	ProjectId   string `json:"projectId" bson:"_id"`
	Html        string `json:"html" bson:"html"`
	Javascript  string `json:"javascript" bson:"javascript"`
	Css         string `json:"css" bson:"css"`
	ProjectName string `json:"projectName" bson:"projectname"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedat"`
}

type UpdateProjResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}