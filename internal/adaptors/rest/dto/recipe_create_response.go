package dto

type RecipeCreateResponse struct {
	RecipeID             string `json:"recipeId"`
	RecipeName           string `json:"recipeName"`
	RecipeDescription    string `json:"recipeDescription"`
	RecipeClassification string `json:"recipeClassification"`
}
