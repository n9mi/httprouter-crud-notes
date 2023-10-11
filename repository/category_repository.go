package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/naomigrain/httprouter-crud-notes/model/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	IsCategoryExistById(ctx context.Context, tx *sql.Tx, id int) bool
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, id int)
}

type CategoryRepositoryImpl struct {
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories where id=?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var category domain.Category
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}

		return category, nil
	}

	return category, errors.New("Category not found")
}

func (r *CategoryRepositoryImpl) IsCategoryExistById(ctx context.Context, tx *sql.Tx, id int) bool {
	SQL := "SELECT 1 FROM categories WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rows.Next()
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (name) VALUES (?)"
	rows, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}

	id, errId := rows.LastInsertId()
	if errId != nil {
		panic(errId)
	}
	category.Id = int(id)

	return category
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	if err != nil {
		panic(err)
	}

	return category
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM categories WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}
