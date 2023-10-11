package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/naomigrain/httprouter-crud-notes/exception"
	"github.com/naomigrain/httprouter-crud-notes/helper"
	"github.com/naomigrain/httprouter-crud-notes/model/domain"
	"github.com/naomigrain/httprouter-crud-notes/model/web"
	"github.com/naomigrain/httprouter-crud-notes/repository"
)

type NoteService interface {
	FindAll(ctx context.Context) []web.NoteShortResponse
	FindById(ctx context.Context, id int) web.NoteResponse
	Create(ctx context.Context, note web.NoteRequest) web.NoteResponse
	Update(ctx context.Context, note web.NoteRequest) web.NoteResponse
	Delete(ctx context.Context, id int)
}

type NoteServiceImpl struct {
	DB                 *sql.DB
	Validate           *validator.Validate
	NoteRepository     repository.NoteRepository
	CategoryRepository repository.CategoryRepository
}

func (s *NoteServiceImpl) FindAll(ctx context.Context) []web.NoteShortResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	notes := s.NoteRepository.FindAll(ctx, tx)
	var notesResponse []web.NoteShortResponse
	for _, n := range notes {
		notesResponse = append(notesResponse, web.NoteShortResponse{
			Id:       n.Id,
			Title:    n.Title,
			Category: n.Category,
		})
	}

	return notesResponse
}

func (s *NoteServiceImpl) FindById(ctx context.Context, id int) web.NoteResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	note, errNotFound := s.NoteRepository.FindById(ctx, tx, id)
	if errNotFound != nil {
		panic(exception.NewNotFoundError("Note not found"))
	}
	noteResponse := web.NoteResponse{
		Id:       note.Id,
		Title:    note.Title,
		Body:     note.Body,
		Category: note.Category,
	}

	return noteResponse
}

func (s *NoteServiceImpl) Create(ctx context.Context, note web.NoteRequest) web.NoteResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	category, errNotFound := s.CategoryRepository.FindById(ctx, tx, note.IdCategory)
	if errNotFound != nil {
		panic(exception.NewNotFoundError("Category not found"))
	}

	errValidate := s.Validate.Struct(note)
	if errValidate != nil {
		panic(errValidate)
	}

	noteDomain := s.NoteRepository.Save(ctx, tx, domain.NoteWithCategoryId{
		Title:      note.Title,
		Body:       note.Body,
		IdCategory: note.IdCategory,
	})
	noteResponse := web.NoteResponse{
		Id:       noteDomain.Id,
		Title:    noteDomain.Title,
		Body:     noteDomain.Body,
		Category: category.Name,
	}

	return noteResponse
}

func (s *NoteServiceImpl) Update(ctx context.Context, note web.NoteRequest) web.NoteResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	if !s.NoteRepository.IsNoteExistById(ctx, tx, note.Id) {
		panic(exception.NewNotFoundError("Note not found"))
	}

	category, errCategory := s.CategoryRepository.FindById(ctx, tx, note.IdCategory)
	if errCategory != nil {
		panic(exception.NewNotFoundError("Category not found"))
	}

	errValidate := s.Validate.Struct(note)
	if errValidate != nil {
		panic(errValidate)
	}

	noteDomain := s.NoteRepository.Update(ctx, tx, domain.NoteWithCategoryId{
		Id:         note.Id,
		Title:      note.Title,
		Body:       note.Body,
		IdCategory: note.IdCategory,
	})
	noteResponse := web.NoteResponse{
		Id:       noteDomain.Id,
		Title:    noteDomain.Title,
		Body:     noteDomain.Body,
		Category: category.Name,
	}

	return noteResponse
}

func (s *NoteServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	if !s.NoteRepository.IsNoteExistById(ctx, tx, id) {
		panic(exception.NewNotFoundError("Note not found"))
	}

	s.NoteRepository.Delete(ctx, tx, id)
}

func NewNoteService(db *sql.DB, validate *validator.Validate,
	noteRepository repository.NoteRepository, categoryRepository repository.CategoryRepository) *NoteServiceImpl {
	return &NoteServiceImpl{
		DB:                 db,
		Validate:           validate,
		NoteRepository:     noteRepository,
		CategoryRepository: categoryRepository,
	}
}
