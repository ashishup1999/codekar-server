package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	UserName    string             `bson:"username"`
	ProjectName string             `bson:"projectname"`
	Html        string             `bson:"html"`
	Javascript  string             `bson:"javascript"`
	Css         string             `bson:"css"`
	Likes       string             `bson:"likes"`
	Comments    string             `bson:"comments"`
	CreatedAt   string             `bson:"createdat"`
	UpdatedAt   string             `bson:"updatedat"`
}

type ProjectMeta struct {
	ProjectId   string `json:"projectId"`
	UserName    string `json:"userName"`
	ProjectName string `json:"projectName"`
	Likes       string `json:"likes"`
	Comments    string `json:"comments"`
}

type AllProjectsResp struct {
	Status   string        `json:"status"`
	Message  string        `json:"message"`
	Projects []ProjectMeta `json:"projects"`
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
