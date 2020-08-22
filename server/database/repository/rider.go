package repository

import (
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Rider interface {
	Save(rider RiderEntity) (string, error)
	FindOne(id string) (RiderEntity, error)
	Update(rider RiderEntity) error
	Delete(id string) error
}

type RiderEntity struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Age              int       `json:"age"`
	Gender           string    `json:"gender"`
	City             string    `json:"city"`
	Cpf              string    `json:"cpf"`
	PaidSubscription bool      `json:"paidSubscription"`
	Sponsors         string    `json:"sponsors"`
	CategoryId       string    `json:"categoryId"`
	CreateAt         time.Time `json:"createAt"`
	UpdateAt         time.Time `json:"updateAt,omitempty"`
}

type RiderRepo struct {
	database database.DbConnection
}

func NewRiderRepository(d database.DbConnection) RiderRepo {
	return RiderRepo{database: d}
}

func (r RiderRepo) Save(rider RiderEntity) (string, error) {

	db := r.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	statement := `INSERT INTO jams.public.RIDERS
				  (rider_id,
				   name, 
				   age, 
				   gender, 
				   city, 
				   cpf, 
				   paid_subscription, 
				   sponsors, 
				   category_id, 
				   created, 
				   updated)
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

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
	statement := `SELECT * FROM jams.public.riders WHERE RIDER_ID=$1`
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

func (r RiderRepo) Delete(id string) error {
	statement := `DELETE FROM jams.public.riders WHERE RIDER_ID=$1`
	db := r.database.ConnectHandle()
	defer db.Close()

	_, err := db.Exec(statement, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r RiderRepo) Update(rider RiderEntity) error {
	statement := `UPDATE jams.public.riders	
				  SET	name=$1,
				      	age=$2, 
				      	gender=$3, 
				      	city=$4, 
				      	cpf=$5, 
				      	paid_subscription=$6, 
				      	sponsors=$7, 
				      	category_id=$8, 
				      	updated=$9 
				  WHERE rider_id=$10`

	db := r.database.ConnectHandle()
	defer db.Close()

	_, err := db.Exec(statement,
		rider.Name,
		rider.Age,
		rider.Gender,
		rider.City,
		rider.Cpf,
		rider.PaidSubscription,
		rider.Sponsors,
		rider.CategoryId,
		rider.UpdateAt,
		rider.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
