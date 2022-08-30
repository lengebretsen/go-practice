package controllers

type ApiError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
