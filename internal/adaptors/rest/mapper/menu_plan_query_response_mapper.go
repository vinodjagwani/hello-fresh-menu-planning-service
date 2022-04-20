package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapMenuToMenuPlanQueryResponse(menu *entity.Menu) *dto.MenuPlanQueryResponse {
	return &dto.MenuPlanQueryResponse{MenuPlanId: menu.MenuID.String(), MenuPlanName: menu.MenuName,
		MenuPlanDescription: menu.MenuDescription, TotalCockingTime: menu.TotalCockingTime, Week: menu.Week, MenuPlanRecipeQueryDetail: MapRecipeToMenuPlanRecipeDetail(menu)}
}

func MapRecipeToMenuPlanRecipeDetail(menu *entity.Menu) []dto.MenuPlanRecipeQueryDetail {
	var recipeDetailList []dto.MenuPlanRecipeQueryDetail
	for _, recipe := range menu.Recipe {
		var ingredientDetailList []dto.MenuPlanIngredientQueryDetail
		for _, ingredient := range recipe.Ingredient {
			ingredientDetailList = append(ingredientDetailList, dto.MenuPlanIngredientQueryDetail{
				IngredientId: ingredient.IngredientID.String(), IngredientName: ingredient.IngredientName,
				IngredientDescription: ingredient.IngredientDescription, IngredientUnit: ingredient.Unit})
		}
		recipeDetailList = append(recipeDetailList,
			dto.MenuPlanRecipeQueryDetail{RecipeID: recipe.RecipeID.String(), RecipeName: recipe.RecipeName, MenuPlanIngredientDetail: ingredientDetailList,
				RecipeDescription: recipe.RecipeDescription, RecipeInstructions: recipe.RecipeInstructions, RecipeClassification: recipe.RecipeClassification})
	}
	return recipeDetailList
}
