package dto

type IngredientCreateResponse struct {
	IngredientId          string `json:"ingredientId"`
	IngredientName        string `json:"ingredientName"`
	IngredientDescription string `json:"ingredientDescription"`
	Unit                  int64  `json:"unit"`
}
