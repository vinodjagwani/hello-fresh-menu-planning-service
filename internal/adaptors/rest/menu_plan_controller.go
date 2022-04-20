package rest

import (
	"github.com/gin-gonic/gin"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/application/ports"
	"hello-fresh-menu-planning-service/internal/application/usecase"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"net/http"
	"strconv"
)

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Create weekly menu plan
// @Schemes
// @Description body
// @Tags MenuPlan
// @Accept json
// @Produce json
// @Param menuPlan body dto.MenuPlanCreateRequest true "Menu plan request body"
// @Success 200 {object} dto.MenuPlanCreateResponse "Menu plan response body"
// @Router /menu-planning-service/api/v1/weekly-menu-plan [post]
func CreateWeeklyMenuPlan(c *gin.Context) {
	var menuCreateRequest dto.MenuPlanCreateRequest
	if err := c.ShouldBindJSON(&menuCreateRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	port := ports.MenuPlanCreatePort{MenuPlanCreateWeeklyUseCase: &usecase.MenuRepositoryForCreateWeeklyMenuPlanUseCase{R: repository.GetMenuRepository(postgres.GetDB())}}
	response, err := port.CreateWeeklyMenuPlan(&menuCreateRequest)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Query weekly menu plan by id
// @Schemes
// @Description body
// @Tags MenuPlan
// @Accept json
// @Produce json
// @Param  menuPlanId  path  int  true "Menu Plan ID"
// @Success 200 {object} dto.MenuPlanCreateResponse "Menu plan response body"
// @Router /menu-planning-service/api/v1/weekly-menu-plan/{menuPlanId} [get]
func FindWeeklyMenuPlanById(c *gin.Context) {
	port := ports.MenuPlanWeeklyQueryUseCase{MenuPlanWeeklyQueryUseCase: &usecase.MenuRepositoryForQueryWeeklyMenuPlanUseCase{R: repository.GetMenuRepository(postgres.GetDB())}}
	response, err := port.FindWeeklyMenuPlanById(c.Param("menuPlanId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &response)
}

// @BasePath /api/v1
// HelloFresh godoc
// @Summary Query all weekly menu plans (from current week and in future)
// @Schemes
// @Description body
// @Tags MenuPlan
// @Produce json
// @Param  page  query  int  false "Page Number"
// @Param  pageSize  query  int  false "Page Size Number"
// @Success 200 {object} []dto.MenuPlanQueryResponse "All Menu plans response body"
// @Router /menu-planning-service/api/v1/weekly-menu-plans [get]
func FindAllWeeklyMenuPlans(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	port := ports.MenuPlanWeeklyQueryUseCase{MenuPlanWeeklyQueryUseCase: &usecase.MenuRepositoryForQueryWeeklyMenuPlanUseCase{R: repository.GetMenuRepository(postgres.GetDB())}}
	response, err := port.FindAllWeeklyMenuPlans(page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, &dto.MenuPlanQueryResponse{})
		return
	}
	c.JSON(http.StatusOK, &response)
}
