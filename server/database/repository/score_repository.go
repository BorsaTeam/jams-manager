package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Score interface {
	Save(score ScoreEntity) (string, error)
}

type ScoreEntity struct {
	Score    float32    `json:"score"`
	Id       string     `json:"id"`
	RiderId  string     `json:"riderId"`
	CreateAt time.Time  `json:"createAt"`
	UpdateAt *time.Time `json:"updateAt,omitempty"`
}

type ScoreRepo struct {
	database database.DbConnection
}

func NewScoreRepository(d database.DbConnection) ScoreRepo {
	return ScoreRepo{database: d}
}

func (s ScoreRepo) Save(score ScoreEntity) (string, error) {

	id, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}
	return id.String(), nil
}
