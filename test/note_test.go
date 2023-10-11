package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/naomigrain/httprouter-crud-notes/model/domain"
	"github.com/naomigrain/httprouter-crud-notes/repository"
	"github.com/stretchr/testify/require"
)

type note struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
}

type noteWithIdCategory struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	IdCategory int    `json:"id_category"`
}

type noteResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   note   `json:"data"`
}

type noteListResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   []note `json:"data"`
}

func createNotes(db *sql.DB, categories []category, notes []note) []note {
	tx, _ := db.Begin()
	noteRepository := repository.NewNoteRepository()

	for i := 0; i < len(categories); i++ {
		noteDomain := noteRepository.Save(context.Background(), tx, domain.NoteWithCategoryId{
			Title:      notes[i].Title,
			Body:       notes[i].Body,
			IdCategory: categories[i].Id,
		})
		tx.Commit()

		notes[i].Id = noteDomain.Id
		notes[i].Category = categories[i].Name

		fmt.Println(notes)
	}

	return notes
}

func TestCreateNotes(t *testing.T) {
	db := setupTestDB()
	categories := createCategories(db, []category{{Name: "Category A"}})
	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/notes"
	title := "Title A"
	body := "Body A"
	requestBody := `{"title": "` + title + `", 
		"body": "` + body + `", "id_category": ` +
		strconv.Itoa(categories[0].Id) + `}`
	request := newTestRequest(url, http.MethodPost, strings.NewReader(requestBody))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseNote noteResponse
	json.Unmarshal(responseBody, &responseNote)

	require.Equal(t, 200, responseNote.Code)
	require.Equal(t, "OK", responseNote.Status)
	require.Equal(t, title, responseNote.Data.Title)
	require.Equal(t, body, responseNote.Data.Body)
	require.Equal(t, categories[0].Name, responseNote.Data.Category)
}

func TestUpdateNotes(t *testing.T) {
	db := setupTestDB()
	categories := createCategories(db, []category{{Name: "Category B"}})
	notes := createNotes(db, categories, []note{{Title: "Note B", Body: "Body B"}})
	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/notes/" + strconv.Itoa(notes[0].Id)
	title := "Title B"
	body := "Body B"
	requestBody := `{"title": "` + title + `", 
		"body": "` + body + `", "id_category": ` +
		strconv.Itoa(categories[0].Id) + `}`
	request := newTestRequest(url, http.MethodPut, strings.NewReader(requestBody))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseNote noteResponse
	json.Unmarshal(responseBody, &responseNote)

	require.Equal(t, 200, responseNote.Code)
	require.Equal(t, "OK", responseNote.Status)
	require.Equal(t, title, responseNote.Data.Title)
	require.Equal(t, body, responseNote.Data.Body)
	require.Equal(t, categories[0].Name, responseNote.Data.Category)
}

func TestDeleteNotes(t *testing.T) {
	db := setupTestDB()
	categories := createCategories(db, []category{{Name: "Category C"}})
	notes := createNotes(db, categories, []note{{Title: "Note C", Body: "Body C"}})
	router := setupTestRouter(db)

	url := "http://127.0.0.1:8000/api/notes/" + strconv.Itoa(notes[0].Id)
	request := newTestRequest(url, http.MethodDelete, strings.NewReader(""))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	responseBody, _ := io.ReadAll(response.Body)
	var responseNote noteResponse
	json.Unmarshal(responseBody, &responseNote)

	require.Equal(t, 200, responseNote.Code)
	require.Equal(t, "OK", responseNote.Status)
	require.Equal(t, "", responseNote.Data.Title)
	require.Equal(t, "", responseNote.Data.Body)
	require.Equal(t, "", responseNote.Data.Category)
}
