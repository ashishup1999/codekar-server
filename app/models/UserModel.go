package models

type User struct {
	UserName  string `bson:"username"`
	Email     string `bson:"email"`
	FullName  string `bson:"fullname"`
	Password  string `bson:"password"`
	CreatedAt string `bson:"createdat"`
	UpdatedAt string `bson:"updatedat"`
}

type CreateUserReq struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type CreateUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
