package model

type ResponseWrapper struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
