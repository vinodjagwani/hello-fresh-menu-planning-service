package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"log"
	"time"
)

type MenuRepository struct {
	con *gorm.DB
}

func GetMenuRepository(con *gorm.DB) *MenuRepository {
	return &MenuRepository{con: con}
}

type IMenuRepository interface {
	SaveMenu(menu *entity.Menu) (*entity.Menu, error)
	FindByMenuId(menuId string) (*entity.Menu, error)
	FindAllWeeksMenus(page int, size int) ([]entity.Menu, error)
}

func (r *MenuRepository) SaveMenu(menu *entity.Menu) (*entity.Menu, error) {
	result := r.con.Create(&menu)
	return menu, result.Error
}

func (r *MenuRepository) FindByMenuId(menuId string) (*entity.Menu, error) {
	var result entity.Menu
	err := r.con.Where(&entity.Menu{MenuID: uuid.Must(uuid.FromString(menuId))}).Preload("Recipe.Ingredient").First(&result).Error
	if err != nil {
		log.Printf("An error occurred while query with menuId: %s with error %s", menuId, err)
	}
	return &result, err
}

func (r *MenuRepository) FindAllWeeksMenus(page int, size int) ([]entity.Menu, error) {
	var result []entity.Menu
	_, week := time.Now().ISOWeek()
	err := r.con.Where("week >= ?", week).Preload("Recipe.Ingredient").Scopes(entity.PaginateScope(page, size)).Find(&result).Error
	if err != nil {
		log.Printf("An error occurred while query all menus with week onwards: %d with error %s", week, err)
	}
	return result, err
}
