package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapRecipeToRecipeUpdateResponse(recipe *entity.Recipe) *dto.RecipeUpdateResponse {
	return &dto.RecipeUpdateResponse{RecipeId: recipe.RecipeID.String(), RecipeName: recipe.RecipeName, RecipeDescription: recipe.RecipeDescription, RecipeClassification: recipe.RecipeClassification, RecipeInstructions: recipe.RecipeInstructions}
}
