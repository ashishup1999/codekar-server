package models

type CompileReq struct {
	Language string `json:"language"`
	Code string `json:"code"`
	Inputs []string `json:"inputs"`
	FileName string `json:"fileName"`	
}

type CompileResp struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Output string `json:"output"`
}