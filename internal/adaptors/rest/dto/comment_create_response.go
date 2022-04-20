package dto

type CommentCreateResponse struct {
	CommentId string `json:"commentId" binding:"required"`
}
