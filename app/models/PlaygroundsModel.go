package models

type Playground struct {
	Id         string `json:"-" bson:"_id,omitempty"`
	UserName   string `json:"userName" bson:"username"`
	PgName     string `json:"pgName" bson:"pgName"`
	Javascript string `json:"javascript" bson:"javascript"`
	Java       string `json:"java" bson:"java"`
	Python     string `json:"python" bson:"python" `
	Cpp        string `json:"cpp" bson:"cpp"`
	Go         string `json:"go" bson:"go"`
	CreatedAt  string `json:"-" bson:"createdat"`
	UpdatedAt  string `json:"-" bson:"updatedat"`
}

type CreatePgReq struct {
	UserName string `json:"userName"`
	PgName   string `json:"pgName"`
}

type CreatePgResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	PgId    string `json:"pgId"`
}

type UpdatePgReq struct {
	PgId       string `json:"pgId" bson:"_id"`
	PgName     string `json:"pgName" bson:"pgName"`
	Javascript string `json:"javascript" bson:"javascript"`
	Java       string `json:"java" bson:"java"`
	Python     string `json:"python" bson:"python" `
	Cpp        string `json:"cpp" bson:"cpp"`
	Go         string `json:"go" bson:"go"`
	UpdatedAt  string `bson:"updatedat"`
}

type UpdatePgResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type PgMeta struct {
	PgId     string `json:"pgId"`
	UserName string `json:"userName"`
	PgName   string `json:"pgName"`
}

type AllPgsResp struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Pgs     []PgMeta `json:"pgs"`
}

type SinglePgResp struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	PgData  Playground `json:"pgData"`
}
