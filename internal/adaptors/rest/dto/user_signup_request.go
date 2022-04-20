package dto

type UserSignUpRequest struct {
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	UserType string `json:"userType" binding:"required"`
	Password string `json:"password" binding:"required"`
}
