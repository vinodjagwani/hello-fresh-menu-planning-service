package dto

type IngredientDetailQueryResponse struct {
	IngredientId          string `json:"ingredientId"`
	IngredientName        string `json:"ingredientName"`
	IngredientDescription string `json:"ingredientDescription"`
	IngredientUnit        int64  `json:"ingredientUnit"`
}
