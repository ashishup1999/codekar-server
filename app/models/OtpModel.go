package models

type OtpReq struct {
	Email string `json:"email" bson:"email"`
	Otp   string `json:"otp" bson:"otp"`
}

type OtpRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
