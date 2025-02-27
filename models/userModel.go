package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type DataRequest struct {
	Message string `json:"message" binding:"required"`
}