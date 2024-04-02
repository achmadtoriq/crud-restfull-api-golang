package service

import (
	"go_api_module/data/request"
	"go_api_module/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
