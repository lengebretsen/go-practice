package controllers

type ApiError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
