package repository

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qulDev/golang-restful-api/helper"
	"github.com/qulDev/golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var category domain.Category

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return domain.Category{}, errors.New("Category is not found")
	}

}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id,name FROM category"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
