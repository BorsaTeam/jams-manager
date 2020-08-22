package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/BorsaTeam/jams-manager/server/database"
)

func TestRiderRepository(t *testing.T) {
	dbConnection := database.NewPgManager()
	dbConnection.TestConnection()

	rr := NewRiderRepository(dbConnection)
	cr := NewCategory(dbConnection)

	categoryId, _ := cr.Save(CategoryEntity{
		Name:     "Pro",
		CreateAt: time.Now(),
		UpdateAt: nil,
	})

	id, _ := rr.Save(
		RiderEntity{
			Name:             "aaa",
			Age:              12,
			Gender:           "aaa",
			City:             "aaa",
			Cpf:              "aa",
			PaidSubscription: false,
			Sponsors:         "aa",
			CategoryId:       categoryId.String(),
			CreateAt:         time.Time{},
			UpdateAt:         time.Time{},
		})

	riderExpect, _ := rr.FindOne(id)
	assert.Equal(t, riderExpect.Id, id)

	_ = rr.Update(RiderEntity{
		Id:               id,
		Name:             "bbb",
		Age:              0,
		Gender:           "",
		City:             "",
		Cpf:              "",
		PaidSubscription: false,
		Sponsors:         "",
		CategoryId:       categoryId.String(),
		CreateAt:         time.Time{},
		UpdateAt:         time.Time{},
	})

	riderUpdated, _ := rr.FindOne(id)
	assert.NotEqual(t, riderExpect.Name, riderUpdated.Name)

	_ = rr.Delete(id)
	_, err := rr.FindOne(id)
	if err == nil {
		t.Error("Find one should return a error")
	}

	assert.Equal(t, err.Error(), "sql: no rows in result set")

}
