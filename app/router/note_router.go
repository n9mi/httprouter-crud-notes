package router

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/naomigrain/httprouter-crud-notes/controller"
	"github.com/naomigrain/httprouter-crud-notes/repository"
	"github.com/naomigrain/httprouter-crud-notes/service"
)

func AddNoteRouter(parentUrl string, router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	subUrl := parentUrl + "notes"

	noteRepository := repository.NewNoteRepository()
	categoryRepository := repository.NewCategoryRepository()
	noteService := service.NewNoteService(db, validate, noteRepository, categoryRepository)
	noteController := controller.NewNoteController(noteService)

	router.GET(subUrl, noteController.GetAll)
	router.GET(subUrl+"/:id", noteController.GetById)
	router.POST(subUrl, noteController.Create)
	router.PUT(subUrl+"/:id", noteController.Update)
	router.DELETE(subUrl+"/:id", noteController.Delete)
}
