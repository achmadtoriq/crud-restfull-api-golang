package repository

import (
	"go_api_module/helper"
	"go_api_module/model"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (n *BookRepositoryImpl) FindAll() []model.Book {
	var book []model.Book
	result := n.Db.Find(&book)
	helper.ErrorPanic(result.Error)

	return book
}
