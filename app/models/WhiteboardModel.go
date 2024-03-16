package models

type Whiteboard struct {
	Id         string `json:"-" bson:"_id,omitempty"`
	UserName   string `json:"userName" bson:"username"`
	WbName     string `json:"wbName" bson:"wbName"`
	Javascript string `json:"javascript" bson:"javascript"`
	Java       string `json:"java" bson:"java"`
	Python     string `json:"python" bson:"python" `
	Cpp        string `json:"cpp" bson:"cpp"`
	Go         string `json:"go" bson:"go"`
	CreatedAt  string `json:"-" bson:"createdat"`
	UpdatedAt  string `json:"-" bson:"updatedat"`
}

type CreateWbReq struct {
	UserName string `json:"userName"`
	WbName   string `json:"wbName"`
}

type CreateWbResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	WbId    string `json:"wbId"`
}

type UpdateWbReq struct {
	WbId       string `json:"wbId" bson:"_id"`
	WbName     string `json:"wbName" bson:"wbName"`
	Javascript string `json:"javascript" bson:"javascript"`
	Java       string `json:"java" bson:"java"`
	Python     string `json:"python" bson:"python" `
	Cpp        string `json:"cpp" bson:"cpp"`
	Go         string `json:"go" bson:"go"`
	UpdatedAt  string `bson:"updatedat"`
}

type UpdateWbResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type WbMeta struct {
	WbId     string `json:"wbId"`
	UserName string `json:"userName"`
	WbName   string `json:"wbName"`
}

type AllWbsResp struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Wbs     []WbMeta `json:"wbs"`
}

type SingleWbResp struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	WbData  Whiteboard `json:"wbData"`
}
