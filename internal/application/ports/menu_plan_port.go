package ports

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/usecase"
)

type MenuPlanCreatePort struct {
	MenuPlanCreateWeeklyUseCase *usecase.MenuRepositoryForCreateWeeklyMenuPlanUseCase
}

type MenuPlanWeeklyQueryUseCase struct {
	MenuPlanWeeklyQueryUseCase *usecase.MenuRepositoryForQueryWeeklyMenuPlanUseCase
}

func (p *MenuPlanCreatePort) CreateWeeklyMenuPlan(request *dto.MenuPlanCreateRequest) (*dto.MenuPlanCreateResponse, error) {
	return p.MenuPlanCreateWeeklyUseCase.CreateWeeklyMenuPlanUseCase(request)
}

func (p *MenuPlanWeeklyQueryUseCase) FindWeeklyMenuPlanById(menuPlanId string) (*dto.MenuPlanQueryResponse, error) {
	return p.MenuPlanWeeklyQueryUseCase.FindWeeklyMenuPlanById(menuPlanId)
}

func (p *MenuPlanWeeklyQueryUseCase) FindAllWeeklyMenuPlans(page int, pageSize int) ([]dto.MenuPlanQueryResponse, error) {
	return p.FindAllWeeklyMenuPlans(page, pageSize)
}
