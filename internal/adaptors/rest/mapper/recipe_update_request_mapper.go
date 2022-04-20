package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapRecipeUpdateRequestToRecipe(request *dto.RecipeUpdateRequest) entity.Recipe {
	return entity.Recipe{RecipeName: request.RecipeName, RecipeDescription: request.RecipeDescription, RecipeClassification: request.RecipeClassification, RecipeInstructions: request.RecipeInstructions}
}
