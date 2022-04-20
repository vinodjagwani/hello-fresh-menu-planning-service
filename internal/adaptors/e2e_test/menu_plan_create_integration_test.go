package e2e_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	_ "gorm.io/driver/postgres"
	"hello-fresh-menu-planning-service/internal/adaptors/rest"
	"hello-fresh-menu-planning-service/internal/adaptors/rest/dto"
	"hello-fresh-menu-planning-service/internal/infra/postgres"
	"hello-fresh-menu-planning-service/internal/infra/repository"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWeeklyMenuPlanSuccess(t *testing.T) {
	postgres.InitPostgresDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	file, _ := ioutil.ReadFile("./mock/menu_plan_create_request.json")
	r.POST("/menu-planning-service/api/v1/weekly-menu-plan", rest.CreateWeeklyMenuPlan)
	req, err := http.NewRequest(http.MethodPost, "/menu-planning-service/api/v1/weekly-menu-plan", bytes.NewBuffer(file))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	var response = dto.MenuPlanCreateResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, response.MenuPlanName, "string")
	postgres.DbConnection.Unscoped().Delete(&entity.Menu{MenuID: uuid.Must(uuid.FromString(response.MenuPlanId))})
}

func TestCreateWeeklyMenuPlanInvalidRequest(t *testing.T) {
	r := gin.Default()
	file, _ := ioutil.ReadFile("./mock/menu_plan_create_invalid_request.json")
	r.POST("/menu-planning-service/api/v1/weekly-menu-plan", rest.CreateWeeklyMenuPlan)
	req, err := http.NewRequest(http.MethodPost, "/menu-planning-service/api/v1/weekly-menu-plan", bytes.NewBuffer(file))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, responseMap["error"], "Key: 'MenuPlanCreateRequest.MenuPlanName' Error:Field validation for 'MenuPlanName' failed on the 'required' tag")
}

func TestCreateWeeklyMenuPlanFindById(t *testing.T) {
	r := gin.Default()
	postgres.InitPostgresDB()
	postgres.DbConnection = postgres.GetDB()
	menuPlan := &entity.Menu{
		MenuID:   uuid.NewV4(),
		MenuName: "test"}
	rep := repository.GetMenuRepository(postgres.DbConnection)
	menu, err := rep.SaveMenu(menuPlan)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	r.GET("/menu-planning-service/api/v1/weekly-menu-plan/:menuPlanId", rest.FindWeeklyMenuPlanById)
	req, err := http.NewRequest(http.MethodGet, "/menu-planning-service/api/v1/weekly-menu-plan/"+menu.MenuID.String(), nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	response := dto.MenuPlanQueryResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, menu.MenuID.String(), response.MenuPlanId)
	postgres.DbConnection.Unscoped().Delete(&menu)
}
