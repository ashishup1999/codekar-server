package models

type User struct {
	Id        string `bson:"_id,omitempty"`
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

type AuthenticateUserReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthenticateUserResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserMata struct {
	UserId   string `json:"userId" bson:"_id,omitempty"`
	UserName string `json:"userName" bson:"username"`
}

type ProfilesResp struct {
	Status   string     `json:"status"`
	Message  string     `json:"message"`
	Profiles []UserMata `json:"profiles"`
}
