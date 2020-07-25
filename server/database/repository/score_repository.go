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
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type ScoreRepo struct {
	database database.DbConnection
}

func NewScoreRepository(d database.DbConnection) ScoreRepo {
	return ScoreRepo{database: d}
}

func (s ScoreRepo) Save(score ScoreEntity) (string, error) {

	statement := `INSERT INTO SCORES
	(SCORE_ID, RIDER_ID, SCORE, CREATED_AT)
	VALUES($1, $2, $3, $4);`

	db := s.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	_, err = db.Exec(statement,
		id,
		score.RiderId,
		score.Score,
		score.CreatedAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
