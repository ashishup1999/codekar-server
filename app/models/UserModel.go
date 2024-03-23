package models

type User struct {
	Id          string   `bson:"_id,omitempty"`
	UserName    string   `bson:"username"`
	Email       string   `bson:"email"`
	FullName    string   `bson:"fullname"`
	Password    string   `bson:"password"`
	Connections []string `bson:"connections"`
	CreatedAt   string   `bson:"createdat"`
	UpdatedAt   string   `bson:"updatedat"`
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

type UserMetaResp struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	UserName string `json:"userName" bson:"username"`
	FullName string `json:"fullName" bson:"fullname"`
}

type ProfilesResp struct {
	Status   string     `json:"status"`
	Message  string     `json:"message"`
	Profiles []UserMata `json:"profiles"`
}

type AddConnectionReq struct {
	UserName1 string `json:"userName1"`
	UserName2 string `json:"userName2"`
}

type StatusResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ConnectionsResponse struct {
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	Connections []string `json:"connections"`
}

type PasswordUpdateReq struct {
	Email    string `json:"email"`
	NewPass string `json:"newPass"`
}
