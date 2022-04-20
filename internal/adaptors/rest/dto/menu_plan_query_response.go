package dto

type MenuPlanQueryResponse struct {
	MenuPlanId                string `json:"menuPlanId"`
	MenuPlanName              string `json:"menuPlanName"`
	MenuPlanDescription       string `json:"menuPlanDescription"`
	TotalCockingTime          string `json:"totalCockingTime"`
	Week                      int    `json:"week"`
	MenuPlanRecipeQueryDetail []MenuPlanRecipeQueryDetail
}

type MenuPlanRecipeQueryDetail struct {
	RecipeID                 string `json:"recipeId"`
	RecipeName               string `json:"recipeName"`
	RecipeDescription        string `json:"recipeDescription"`
	RecipeInstructions       string `json:"recipeInstructions"`
	RecipeClassification     string `json:"recipeClassification"`
	MenuPlanIngredientDetail []MenuPlanIngredientQueryDetail
}

type MenuPlanIngredientQueryDetail struct {
	IngredientId          string `json:"ingredientId"`
	IngredientName        string `json:"ingredientName"`
	IngredientDescription string `json:"ingredientDescription"`
	IngredientUnit        int64  `json:"ingredientUnit"`
}
