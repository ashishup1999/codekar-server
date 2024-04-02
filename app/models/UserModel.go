package models

type User struct {
	Id            string   `bson:"_id,omitempty"`
	UserName      string   `bson:"username"`
	Email         string   `bson:"email"`
	FullName      string   `bson:"fullname"`
	Password      string   `bson:"password"`
	Connections   []string `bson:"connections"`
	ConnectionReq []string `bson:"connReqs"`
	CreatedAt     string   `bson:"createdat"`
	UpdatedAt     string   `bson:"updatedat"`
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
	Status      string   `json:"status"`
	Message     string   `json:"message"`
	UserName    string   `json:"userName" bson:"username"`
	FullName    string   `json:"fullName" bson:"fullname"`
	Connections []string `json:"connections" bson:"connections"`
}

type ProfilesResp struct {
	Status   string     `json:"status"`
	Message  string     `json:"message"`
	Profiles []UserMata `json:"profiles"`
}

type AddConnectionReq struct {
	Sender   string `json:"sender"`
	Reciever string `json:"reciever"`
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

type AllConnReqsResp struct {
	Status        string   `json:"status"`
	Message       string   `json:"message"`
	ConnectionReq []string `json:"connectionReqs"`
}

type PasswordUpdateReq struct {
	Email   string `json:"email"`
	NewPass string `json:"newPass"`
}

type EditUserReq struct {
	UserName     string `json:"userName"`
	NewUserName  string `json:"newUserName"`
	NewEmail     string `json:"newEmail"`
	NewFullName  string `json:"newFullName"`
	CurrPassword string `json:"currPassword"`
	NewPassword  string `json:"newPassword"`
}
