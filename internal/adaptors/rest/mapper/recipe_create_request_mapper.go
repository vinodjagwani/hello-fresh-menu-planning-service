package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapRecipeCreateRequestToRecipe(request *dto.RecipeCreateRequest) entity.Recipe {
	return entity.Recipe{RecipeName: request.RecipeName, RecipeDescription: request.RecipeDescription, RecipeInstructions: request.RecipeInstructions,
		RecipeClassification: request.RecipeClassification}
}
