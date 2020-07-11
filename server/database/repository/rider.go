package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Rider interface {
	Save(rider RiderEntity) (string, error)
	FindOne(id string) (RiderEntity, error)
}

type RiderEntity struct {
	Id               string     `json:"id"`
	Name             string     `json:"name"`
	Age              int        `json:"age"`
	Gender           string     `json:"gender"`
	City             string     `json:"city"`
	Cpf              string     `json:"cpf"`
	PaidSubscription bool       `json:"paidSubscription"`
	Sponsors         string     `json:"sponsors"`
	CategoryId       string     `json:"categoryId"`
	CreateAt         time.Time  `json:"createAt"`
	UpdateAt         *time.Time `json:"updateAt,omitempty"`
}

type RiderRepo struct {
	database database.DbConnection
}

func NewRiderRepository(d database.DbConnection) RiderRepo {
	return RiderRepo{database: d}
}

func (r RiderRepo) Save(rider RiderEntity) (string, error) {

	statement := `INSERT INTO public.RIDERS
				  (rider_id, name, age, gender, city, cpf, paid_subscription, sponsors, category_id, created, updated)
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

	db := r.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	_, err = db.Exec(statement,
		id,
		rider.Name,
		rider.Age,
		rider.Gender,
		rider.City,
		rider.Cpf,
		rider.PaidSubscription,
		rider.Sponsors,
		rider.CategoryId,
		rider.CreateAt,
		rider.UpdateAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
func (r RiderRepo) FindOne(id string) (RiderEntity, error) {
	statement := `SELECT * FROM RIDERS WHERE RIDER_ID=$1`
	db := r.database.ConnectHandle()
	defer db.Close()

	re := RiderEntity{}

	row := db.QueryRow(statement, id)
	if err := row.Scan(
		&re.Id,
		&re.Name,
		&re.Age,
		&re.Gender,
		&re.City,
		&re.Cpf,
		&re.PaidSubscription,
		&re.Sponsors,
		&re.CategoryId,
		&re.CreateAt,
		&re.UpdateAt,
	); err != nil {
		return RiderEntity{}, err
	}
	return re, nil
}
