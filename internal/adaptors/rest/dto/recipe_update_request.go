package dto

type RecipeUpdateRequest struct {
	RecipeName           string `json:"recipeName" binding:"required"`
	RecipeDescription    string `json:"recipeDescription" binding:"required"`
	RecipeClassification string `json:"recipeClassification" binding:"required"`
	RecipeInstructions   string `json:"recipeInstructions" binding:"required"`
}
