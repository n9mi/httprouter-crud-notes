package router

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/naomigrain/httprouter-crud-notes/controller"
	"github.com/naomigrain/httprouter-crud-notes/repository"
	"github.com/naomigrain/httprouter-crud-notes/service"
)

func AddCategoryRouter(parentUrl string, router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	subUrl := parentUrl + "categories"

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, validate, categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router.GET(subUrl, categoryController.GetAll)
	router.GET(subUrl+"/:id", categoryController.GetById)
	router.POST(subUrl, categoryController.Create)
	router.PUT(subUrl+"/:id", categoryController.Update)
	router.DELETE(subUrl+"/:id", categoryController.Delete)
}
