package mapper

import (
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

func MapMenuToMenuCreateResponse(menu *entity.Menu) *dto.MenuPlanCreateResponse {
	return &dto.MenuPlanCreateResponse{MenuPlanName: menu.MenuName, MenuPlanId: menu.MenuID.String()}
}
