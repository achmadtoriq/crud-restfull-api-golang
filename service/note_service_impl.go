package service

import (
	"github.com/go-playground/validator/v10"
	"go_api_module/data/request"
	"go_api_module/data/response"
	"go_api_module/helper"
	"go_api_module/model"
	"go_api_module/repository"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	validate       *validator.Validate
}

func NewNoteServiceImpl(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		validate:       validate,
	}
}

func (n *NoteServiceImpl) Create(note request.CreateNoteRequest) {
	//TODO implement me
	err := n.validate.Struct(note)
	helper.ErrorPanic(err)
	noteModel := model.Note{
		Content: note.Content,
	}

	n.NoteRepository.Save(noteModel)
}

func (n *NoteServiceImpl) Update(note request.UpdateNoteRequest) {
	//TODO implement me
	noteData, err := n.NoteRepository.FindById(note.Id)
	helper.ErrorPanic(err)
	noteData.Content = note.Content
	n.NoteRepository.Update(noteData)
}

func (n *NoteServiceImpl) Delete(noteId int) {
	//TODO implement me
	n.NoteRepository.Delete(noteId)
}

func (n *NoteServiceImpl) FindById(noteId int) response.NoteResponse {
	//TODO implement me
	noteData, err := n.NoteRepository.FindById(noteId)
	helper.ErrorPanic(err)
	noteResponse := response.NoteResponse{
		Id:      noteData.Id,
		Content: noteData.Content,
	}

	return noteResponse
}

func (n *NoteServiceImpl) FindAll() []response.NoteResponse {
	//TODO implement me
	result := n.NoteRepository.FindAll()
	var notes []response.NoteResponse

	for _, value := range result {
		note := response.NoteResponse{Id: value.Id, Content: value.Content}
		notes = append(notes, note)
	}

	return notes
}
