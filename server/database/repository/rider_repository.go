package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Rider interface {
	Save(rider RiderEntity) (string, error)
}

type RiderEntity struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Age              int        `json:"age"`
	Gender           string     `json:"gender"`
	City             string     `json:"city"`
	Cpf              string     `json:"cpf"`
	PaidSubscription bool       `json:"paidSubscription"`
	Sponsors         []string   `json:"sponsors"`
	CategoryId       string     `json:"categoryId"`
	CreateAt         time.Time  `json:"createAt"`
	UpdateAt         *time.Time `json:"updateAt,omitempty"`
}

type rider struct {
	database database.DbConnection
}

func NewRiderRepository(d database.DbConnection) *rider {
	return &rider{database: d}
}

func (r *rider) Save(rider RiderEntity) (string, error) {
	statement := `INSERT INTO public.rider
				  (rider_id, name, age, gender, city, cpf, paidsubscription, sponsors, category_id, created)
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);`

	db := r.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	_, err = db.Exec(statement, id, rider.Age, rider.Gender, rider.City, rider.Cpf, rider.PaidSubscription, rider.Sponsors, rider.CategoryId, rider.CreateAt)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return id.String(), nil
}
