package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapMenuCreateRequestToMenu(request *dto.MenuPlanCreateRequest) entity.Menu {
	return entity.Menu{MenuName: request.MenuPlanName, MenuDescription: request.MenuPlanDescription, TotalCockingTime: request.TotalCookingTime, Week: request.Week, Recipe: MapMenuPlanRecipeCreateRequestToRecipe(request)}
}

func MapMenuPlanRecipeCreateRequestToRecipe(request *dto.MenuPlanCreateRequest) []entity.Recipe {
	var recipeList []entity.Recipe
	for _, value := range request.MenuPlanRecipeCreateRequest {
		var ingredientList []entity.Ingredient
		for _, ingredientValue := range value.MenuPlanIngredientCreateRequest {
			ingredientList = append(ingredientList,
				entity.Ingredient{IngredientName: ingredientValue.IngredientName,
					IngredientDescription: ingredientValue.IngredientDescription,
					Unit:                  ingredientValue.IngredientUnit,
				})
		}
		recipeList = append(recipeList,
			entity.Recipe{RecipeName: value.RecipeName,
				RecipeDescription: value.RecipeDescription,
				Ingredient:        ingredientList})
	}
	return recipeList
}
