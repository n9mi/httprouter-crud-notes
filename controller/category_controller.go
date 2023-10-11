package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/naomigrain/httprouter-crud-notes/exception"
	"github.com/naomigrain/httprouter-crud-notes/helper"
	"github.com/naomigrain/httprouter-crud-notes/model/web"
	"github.com/naomigrain/httprouter-crud-notes/service"
)

type CategoryController interface {
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func (c *CategoryControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categories := c.Service.FindAll(r.Context())

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categories,
	}

	helper.EncodeJson(w, response)
}

func (c *CategoryControllerImpl) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Category not found"))
	}

	category := c.Service.FindById(r.Context(), idInt)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   category,
	}

	helper.EncodeJson(w, response)
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var categoryRequest web.CategoryRequest
	helper.DecodeJson(r, &categoryRequest)

	categoryResponse := c.Service.Create(r.Context(), categoryRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.EncodeJson(w, response)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var categoryRequest web.CategoryRequest
	helper.DecodeJson(r, &categoryRequest)

	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Category not found"))
	}

	categoryRequest.Id = idInt
	categoryResponse := c.Service.Update(r.Context(), categoryRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.EncodeJson(w, response)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Category not found"))
	}

	c.Service.Delete(r.Context(), idInt)

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.EncodeJson(w, response)
}

func NewCategoryController(service service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		Service: service,
	}
}
