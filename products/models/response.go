package models

type SuccessResponse struct {
	Response string      `json:"response"`
	Data     interface{} `json:"data"`
}

type ErrorResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
