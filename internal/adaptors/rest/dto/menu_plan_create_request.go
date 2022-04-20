package dto

type MenuPlanCreateRequest struct {
	MenuPlanName                string                        `json:"menuPlanName" binding:"required"`
	MenuPlanDescription         string                        `json:"menuPlanDescription"`
	Week                        int                           `json:"week"`
	TotalCookingTime            string                        `json:"totalCockingTime"`
	MenuPlanRecipeCreateRequest []MenuPlanRecipeCreateRequest `json:"menuPlanRecipeCreateRequest"`
}

type MenuPlanRecipeCreateRequest struct {
	RecipeName                      string                            `json:"recipeName" binding:"required"`
	RecipeDescription               string                            `json:"recipeDescription"`
	RecipeClassification            string                            `json:"recipeClassification"`
	MenuPlanIngredientCreateRequest []MenuPlanIngredientCreateRequest `json:"menuPlanIngredientCreateRequest"`
}

type MenuPlanIngredientCreateRequest struct {
	IngredientName        string `json:"ingredientName" binding:"required"`
	IngredientDescription string `json:"ingredientDescription" binding:"required"`
	IngredientUnit        int64  `json:"ingredientUnit"`
}
