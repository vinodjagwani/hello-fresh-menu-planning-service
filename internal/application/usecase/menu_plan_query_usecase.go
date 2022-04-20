package usecase

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/mapper"
	"hello-fresh-menu-planning-service/internal/infra/repository"
)

type MenuRepositoryForQueryWeeklyMenuPlanUseCase struct {
	R *repository.MenuRepository
}

type IMenuPlanWeeklyQueryUseCase interface {
	FindWeeklyMenuPlanById(menuPlanId string) (*dto.MenuPlanQueryResponse, error)
	FindAllWeeklyMenuPlans(page int, pageSize int) (*[]dto.MenuPlanQueryResponse, error)
}

func (uc *MenuRepositoryForQueryWeeklyMenuPlanUseCase) FindWeeklyMenuPlanById(menuPlanId string) (*dto.MenuPlanQueryResponse, error) {
	menuPlan, err := uc.R.FindByMenuId(menuPlanId)
	if err != nil {
		return &dto.MenuPlanQueryResponse{}, err
	}
	return mapper.MapMenuToMenuPlanQueryResponse(menuPlan), err
}

func (uc *MenuRepositoryForQueryWeeklyMenuPlanUseCase) FindAllWeeklyMenuPlans(page int, pageSize int) ([]*dto.MenuPlanQueryResponse, error) {
	menuPlans, err := uc.R.FindAllWeeksMenus(page, pageSize)
	var response = []*dto.MenuPlanQueryResponse{}
	if err != nil {
		return []*dto.MenuPlanQueryResponse{}, err
	}
	for _, menuPlan := range menuPlans {
		response = append(response, mapper.MapMenuToMenuPlanQueryResponse(&menuPlan))
	}
	return response, err
}
