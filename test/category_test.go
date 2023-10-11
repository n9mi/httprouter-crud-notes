package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naomigrain/httprouter-crud-notes/model/domain"
	"github.com/naomigrain/httprouter-crud-notes/repository"
	"github.com/stretchr/testify/require"
)

// go test -v ./test

type categoryResponse struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   category `json:"data"`
}

type categoryListResponse struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Data   []category `json:"data"`
}

type category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func createCategories(db *sql.DB, categories []category) []category {
	categoryRepository := repository.NewCategoryRepository()

	for i, c := range categories {
		tx, _ := db.Begin()
		cDomain := categoryRepository.Save(context.Background(), tx, domain.Category{
			Id:   c.Id,
			Name: c.Name,
		})
		tx.Commit()
		categories[i].Id = cDomain.Id
	}

	return categories
}

func TestCreateCategory(t *testing.T) {
	db := setupTestDB()
	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/categories"
	categoryName := "Test Category"
	requestBody := `{"name": "` + categoryName + `"}`
	request := newTestRequest(url, http.MethodPost, strings.NewReader(requestBody))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseCategory categoryResponse
	json.Unmarshal(responseBody, &responseCategory)

	require.Equal(t, 200, responseCategory.Code)
	require.Equal(t, "OK", responseCategory.Status)
	require.Equal(t, categoryName, responseCategory.Data.Name)
}

func TestUpdateCategory(t *testing.T) {
	db := setupTestDB()
	categories := []category{
		{Name: "A"},
	}
	categories = createCategories(db, categories)

	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/categories/" + strconv.Itoa(categories[0].Id)

	categoryName := "Test A"
	requestBody := `{"name": "` + categoryName + `"}`
	request := newTestRequest(url, http.MethodPut, strings.NewReader(requestBody))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseCategory categoryResponse
	json.Unmarshal(responseBody, &responseCategory)

	require.Equal(t, 200, responseCategory.Code)
	require.Equal(t, "OK", responseCategory.Status)
	require.Equal(t, categoryName, responseCategory.Data.Name)
}

func TestDeleteCategory(t *testing.T) {
	db := setupTestDB()
	categories := []category{
		{Name: "A"},
	}
	categories = createCategories(db, categories)

	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/categories/" + strconv.Itoa(categories[0].Id)

	request := newTestRequest(url, http.MethodDelete, strings.NewReader(""))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseCategory categoryResponse
	json.Unmarshal(responseBody, &responseCategory)

	require.Equal(t, 200, responseCategory.Code)
	require.Equal(t, "OK", responseCategory.Status)
	require.Equal(t, "", responseCategory.Data.Name)
}

func TestGetAllCategory(t *testing.T) {
	db := setupTestDB()
	truncateTestDB(db)
	categories := []category{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
	}
	categories = createCategories(db, categories)

	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/categories"

	request := newTestRequest(url, http.MethodGet, strings.NewReader(""))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var categoryListResponse categoryListResponse
	json.Unmarshal(responseBody, &categoryListResponse)

	require.Equal(t, 200, categoryListResponse.Code)
	require.Equal(t, "OK", categoryListResponse.Status)
	require.Equal(t, len(categories), len(categoryListResponse.Data))

	for i := 0; i < len(categories); i++ {
		require.Equal(t, categories[i].Name, categoryListResponse.Data[i].Name)
	}
}

func TestGetByIdCategory(t *testing.T) {
	db := setupTestDB()

	categories := []category{
		{Name: "C"},
	}
	categories = createCategories(db, categories)

	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/categories/" + strconv.Itoa(categories[0].Id)

	request := newTestRequest(url, http.MethodGet, strings.NewReader(""))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseCategory categoryResponse
	json.Unmarshal(responseBody, &responseCategory)

	require.Equal(t, 200, responseCategory.Code)
	require.Equal(t, "OK", responseCategory.Status)
	require.Equal(t, categories[0].Name, responseCategory.Data.Name)
}
