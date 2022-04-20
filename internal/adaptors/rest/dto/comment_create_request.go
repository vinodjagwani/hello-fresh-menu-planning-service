package dto

type CommentCreateRequest struct {
	Comment string `json:"comment" binding:"required"`
	UserID  string `json:"userID" binding:"required"`
	MenuPlanId  string `json:"menuPlanId" binding:"required"`
}
