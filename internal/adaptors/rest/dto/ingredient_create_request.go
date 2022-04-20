package dto

type IngredientCreateRequest struct {
	IngredientName        string `json:"ingredientName" binding:"required"`
	IngredientDescription string `json:"ingredientDescription" binding:"required"`
	IngredientUnit        int64  `json:"ingredientUnit" binding:"required"`
}
