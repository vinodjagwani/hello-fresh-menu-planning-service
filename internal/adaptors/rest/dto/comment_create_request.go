package dto

type CommentCreateRequest struct {
	Comment string `json:"comment" binding:"required"`
	UserID  string `json:"userID" binding:"required"`
}
