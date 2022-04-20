package repository

import (
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"log"
)

type UserRepository struct {
	con *gorm.DB
}

func GetUserRepository(con *gorm.DB) *UserRepository {
	return &UserRepository{con: con}
}

type IUserRepository interface {
	SaveUser(user *entity.User) (*entity.User, error)
	FindUserByEmailAndPassword(email string, password string) (*entity.User, error)
}

func (r *UserRepository) SaveUser(user *entity.User) (*entity.User, error) {
	result := r.con.Create(&user)
	return user, result.Error
}

func (r *UserRepository) FindUserByEmailAndPassword(email string, password string) (*entity.User, error) {
	var result entity.User
	err := r.con.Where(&entity.User{Email: email, Password: password}).Find(&result).Error
	if err != nil {
		log.Printf("An error occurred while query with email: %s with error %s", email, err)
	}
	return &result, err
}
