package ports

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/usecase"
)

type RecipeCreatePort struct {
	RecipeCreateUseCase *usecase.RecipeCreateUseCase
}

type RecipeDeletePort struct {
	RecipeDeleteUseCase *usecase.RecipeDeleteUseCase
}

type RecipeQueryPort struct {
	RecipeQueryUseCase *usecase.RecipeQueryUseCase
}

type RecipeUpdatePort struct {
	RecipeUpdateUseCase *usecase.RecipeUpdateUseCase
}

func (p *RecipeCreatePort) CreateRecipe(request *dto.RecipeCreateRequest) (*dto.RecipeCreateResponse, error) {
	return p.RecipeCreateUseCase.CreateRecipe(request)
}

func (p *RecipeUpdatePort) UpdateRecipe(recipeId string, request *dto.RecipeUpdateRequest) (*dto.RecipeUpdateResponse, error) {
	return p.RecipeUpdateUseCase.UpdateRecipe(recipeId, request)
}

func (p *RecipeQueryPort) FindRecipeById(recipeId string) (*dto.RecipeDetailQueryResponse, error) {
	return p.RecipeQueryUseCase.FindRecipeById(recipeId)
}

func (p *RecipeDeletePort) DeleteRecipeById(recipeId string) error {
	return p.RecipeDeleteUseCase.DeleteRecipeById(recipeId)
}
