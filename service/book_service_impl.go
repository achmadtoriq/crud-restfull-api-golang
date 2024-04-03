package service

import (
	"go_api_module/data/response"
	"go_api_module/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (n *BookServiceImpl) FindAll() []response.BookResponse {
	//TODO implement me
	result := n.BookRepository.FindAll()
	var books []response.BookResponse

	for _, value := range result {
		book := response.BookResponse{Id: value.Id, Title: value.Title}
		books = append(books, book)
	}

	return books
}
