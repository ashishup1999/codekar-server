package models

type ProfilesByNameReq struct {
	Name string `json:"name"`
	PageNo int64 `json:"pageNo"`
}

type ProjectsByNameReq struct {
	ProjName string `json:"projName"`
	PageNo int64 `json:"pageNo"`
}

type PgsByNameReq struct {
	PgName string `json:"pgName"`
	PageNo int64 `json:"pageNo"`
}

type WbsByNameReq struct {
	WbName string `json:"wbName"`
	PageNo int64 `json:"pageNo"`
}