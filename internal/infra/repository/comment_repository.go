package repository

import (
	"fmt"
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
)

type CommentRepository struct {
	con *gorm.DB
}

func GetCommentRepository(con *gorm.DB) *CommentRepository {
	return &CommentRepository{con: con}
}

type ICommentRepository interface {
	SaveComment(comment *entity.Comment) (*entity.Comment, error)
}

func (r *CommentRepository) SaveComment(comment *entity.Comment) (*entity.Comment, error) {
	fmt.Println(comment)
	result := r.con.Create(&comment)
	return comment, result.Error
}
