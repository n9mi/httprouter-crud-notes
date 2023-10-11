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

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, id int) web.CategoryResponse
	Create(ctx context.Context, category web.CategoryRequest) web.CategoryResponse
	Update(ctx context.Context, category web.CategoryRequest) web.CategoryResponse
	Delete(ctx context.Context, id int)
}

type CategoryServiceImpl struct {
	DB         *sql.DB
	Validate   *validator.Validate
	Repository repository.CategoryRepository
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	categories := s.Repository.FindAll(ctx, tx)
	var categoriesResponse []web.CategoryResponse
	for _, c := range categories {
		categoryResponse := web.CategoryResponse{
			Id:   c.Id,
			Name: c.Name,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, id int) web.CategoryResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}

	category, errNotFound := s.Repository.FindById(ctx, tx, id)
	if errNotFound != nil {
		panic(exception.NewNotFoundError(errNotFound.Error()))
	}

	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, category web.CategoryRequest) web.CategoryResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	errValidate := s.Validate.Struct(category)
	if errValidate != nil {
		panic(errValidate)
	}

	categoryDomain := s.Repository.Save(ctx, tx, domain.Category{
		Name: category.Name,
	})
	categoryResponse := web.CategoryResponse{
		Id:   categoryDomain.Id,
		Name: categoryDomain.Name,
	}

	return categoryResponse
}

func (s *CategoryServiceImpl) Update(ctx context.Context, category web.CategoryRequest) web.CategoryResponse {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	if !s.Repository.IsCategoryExistById(ctx, tx, category.Id) {
		panic(exception.NewNotFoundError("Category not found"))
	}

	errValidate := s.Validate.Struct(category)
	if errValidate != nil {
		panic(errValidate)
	}

	categoryDomain := s.Repository.Update(ctx, tx, domain.Category{
		Id:   category.Id,
		Name: category.Name,
	})

	categoryResponse := web.CategoryResponse{
		Id:   categoryDomain.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	if !s.Repository.IsCategoryExistById(ctx, tx, id) {
		panic(exception.NewNotFoundError("Category not found"))
	}

	s.Repository.Delete(ctx, tx, id)
}

func NewCategoryService(db *sql.DB, validate *validator.Validate, repository repository.CategoryRepository) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:         db,
		Validate:   validate,
		Repository: repository,
	}
}
