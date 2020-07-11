package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Category interface {
	Save(category CategoryEntity) (CategoryId, error)
	FindAll() ([]CategoryEntity, error)
	Delete(id CategoryId) error
}

type CategoryId string

func (c CategoryId) String() string {
	return string(c)
}

type CategoryEntity struct {
	Id       CategoryId
	Name     string
	CreateAt time.Time
	UpdateAt *time.Time
}

type CategoryRepo struct {
	db database.DbConnection
}

func NewCategory(db database.DbConnection) CategoryRepo {
	return CategoryRepo{db: db}
}

func (ca CategoryRepo) Save(category CategoryEntity) (CategoryId, error) {
	stmt := `INSERT INTO categories VALUES($1, $2, $3)`

	db := ca.db.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	if _, err := db.Exec(stmt, id, category.Name, category.CreateAt); err != nil {
		return "", err
	}

	return CategoryId(id.String()), nil
}

func (ca CategoryRepo) FindAll() ([]CategoryEntity, error) {
	stmt := `SELECT * FROM categories`

	db := ca.db.ConnectHandle()
	defer db.Close()

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	var categories []CategoryEntity
	for rows.Next() {
		category := CategoryEntity{}
		if err := rows.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (ca CategoryRepo) Delete(id CategoryId) error {
	stmt := `DELETE FROM categories WHERE category_id = $1`

	db := ca.db.ConnectHandle()
	defer db.Close()

	if _, err := db.Exec(stmt, id); err != nil {
		return err
	}

	return nil
}
