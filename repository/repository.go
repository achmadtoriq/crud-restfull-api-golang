package repository

import "go_api_module/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(noteId int)
	FindById(noteId int) (model.Note, error)
	FindAll() []model.Note
}

type BookRepository interface {
	FindAll() []model.Book
}
