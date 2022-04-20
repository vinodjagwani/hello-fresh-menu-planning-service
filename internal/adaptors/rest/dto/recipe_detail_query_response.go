package dto

type RecipeDetailQueryResponse struct {
	RecipeId             string `json:"recipeId" `
	RecipeName           string `json:"recipeName"`
	RecipeDescription    string `json:"recipeDescription"`
	RecipeClassification string `json:"recipeClassification"`
}
