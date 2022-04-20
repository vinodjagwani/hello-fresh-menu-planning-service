package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type MenuRepositoryForCreateWeeklyMenuPlanUseCase struct {
	R *repository.MenuRepository
}

type ICreateWeeklyMenuPlanUseCase interface {
	CreateWeeklyMenuPlanUseCase(request *dto.MenuPlanCreateRequest) (*dto.MenuPlanCreateResponse, error)
}

func (uc *MenuRepositoryForCreateWeeklyMenuPlanUseCase) CreateWeeklyMenuPlanUseCase(request *dto.MenuPlanCreateRequest) (*dto.MenuPlanCreateResponse, error) {
	var menu = mapper.MapMenuCreateRequestToMenu(request)
	savedMenu, err := uc.R.SaveMenu(&menu)
	if err != nil {
		return mapper.MapMenuToMenuCreateResponse(nil), err
	}
	return mapper.MapMenuToMenuCreateResponse(savedMenu), err
}
