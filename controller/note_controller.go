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

type NoteController interface {
	GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type NoteControllerImpl struct {
	Service service.NoteService
}

func (c *NoteControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	notes := c.Service.FindAll(r.Context())
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   notes,
	}

	helper.EncodeJson(w, response)
}

func (c *NoteControllerImpl) GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Note not found"))
	}

	note := c.Service.FindById(r.Context(), idInt)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   note,
	}

	helper.EncodeJson(w, response)
}

func (c *NoteControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var noteRequest web.NoteRequest
	helper.DecodeJson(r, &noteRequest)

	noteResponse := c.Service.Create(r.Context(), noteRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   noteResponse,
	}

	helper.EncodeJson(w, response)
}

func (c *NoteControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Note not found"))
	}

	var noteRequest web.NoteRequest
	helper.DecodeJson(r, &noteRequest)
	noteRequest.Id = idInt

	noteResponse := c.Service.Update(r.Context(), noteRequest)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   noteResponse,
	}

	helper.EncodeJson(w, response)
}

func (c *NoteControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		panic(exception.NewNotFoundError("Note not found"))
	}

	c.Service.Delete(r.Context(), idInt)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.EncodeJson(w, response)
}

func NewNoteController(service service.NoteService) *NoteControllerImpl {
	return &NoteControllerImpl{
		Service: service,
	}
}
