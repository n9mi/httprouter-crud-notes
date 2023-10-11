package router

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	parentUrl := "/api/"
	router := httprouter.New()

	AddCategoryRouter(parentUrl, router, db, validate)
	AddNoteRouter(parentUrl, router, db, validate)

	return router
}
