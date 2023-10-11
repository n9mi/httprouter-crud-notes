package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/naomigrain/httprouter-crud-notes/model/domain"
)

type NoteRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Note
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Note, error)
	IsNoteExistById(ctx context.Context, tx *sql.Tx, id int) bool
	Save(ctx context.Context, tx *sql.Tx, note domain.NoteWithCategoryId) domain.NoteWithCategoryId
	Update(ctx context.Context, tx *sql.Tx, note domain.NoteWithCategoryId) domain.NoteWithCategoryId
	Delete(ctx context.Context, tx *sql.Tx, id int)
}

type NoteRepositoryImpl struct {
}

func (r *NoteRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Note {
	SQL := `SELECT notes.id, notes.title, categories.name as category
		FROM notes INNER JOIN categories where notes.id_category = categories.id`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var notes []domain.Note
	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Category)
		if err != nil {
			panic(err)
		}
		notes = append(notes, note)
	}

	return notes
}

func (r *NoteRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Note, error) {
	SQL := `SELECT notes.id, notes.title, notes.body, categories.name as category
		FROM notes INNER JOIN categories where notes.id_category = categories.id
		AND notes.id = ?`
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var note domain.Note
	if rows.Next() {
		err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.Category)
		if err != nil {
			panic(err)
		}

		return note, nil
	}

	return note, errors.New("Category not found")
}

func (r *NoteRepositoryImpl) IsNoteExistById(ctx context.Context, tx *sql.Tx, id int) bool {
	SQL := `SELECT 1 FROM notes where id=?`
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows.Next()
}

func (r *NoteRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, note domain.NoteWithCategoryId) domain.NoteWithCategoryId {
	SQL := `INSERT INTO notes (title, body, id_category) VALUES (?, ?, ?)`
	rows, err := tx.ExecContext(ctx, SQL, note.Title, note.Body, note.IdCategory)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

	id, errId := rows.LastInsertId()
	if errId != nil {
		panic(errId)
	}
	note.Id = int(id)

	return note
}

func (r *NoteRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, note domain.NoteWithCategoryId) domain.NoteWithCategoryId {
	SQL := `UPDATE notes SET title = ?, body = ?, id_category = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, note.Title, note.Body, note.IdCategory, note.Id)
	if err != nil {
		panic(err)
	}

	return note
}

func (r *NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := `DELETE FROM notes WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}

func NewNoteRepository() *NoteRepositoryImpl {
	return &NoteRepositoryImpl{}
}
