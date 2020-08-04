package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server/database"
)

type Judge interface {
	Save(judge JudgeEntity) (string, error)
}

type JudgeEntity struct {
	JudgeId   string     `json:"judgeId"`
	Password  float64    `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type JudgeRepo struct {
	database database.DbConnection
}

func NewJudgeRepository(d database.DbConnection) JudgeRepo {
	return JudgeRepo{database: d}
}

func (j JudgeRepo) save(judge JudgeEntity) (string, error) {
	statement := `INSERT INTO JUDGE
	(JUDGE_ID, PASSWORD, CREATED_AT)
	VALUES($1, $2, $3);`

	db := j.database.ConnectHandle()
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	_, err = db.Exec(statement,
		judge.JudgeId,
		judge.Password,
		judge.CreatedAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil

}

