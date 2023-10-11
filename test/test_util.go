package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/naomigrain/httprouter-crud-notes/app/database"
	"github.com/naomigrain/httprouter-crud-notes/app/router"
	"github.com/naomigrain/httprouter-crud-notes/config"
)

func setupTestDB() *sql.DB {
	testConfig := config.GetTestDBConfig(true)
	db := database.NewDB(testConfig)

	return db
}

func setupTestRouter(db *sql.DB) *httprouter.Router {
	validate := validator.New()
	router := router.NewRouter(db, validate)

	return router
}

func truncateTestDB(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE categories;")
	db.Exec("TRUNCATE TABLE notes;")
}

func newTestRequest(url string, method string, requestBody *strings.Reader) *http.Request {
	request := httptest.NewRequest(method, url, requestBody)
	request.Header.Add("Content-Type", "application/json")

	return request
}
