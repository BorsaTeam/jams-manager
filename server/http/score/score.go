package score

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/postgres/repository"
)

type Manager struct {
	scoreRepository repository.Score
}

func NewHandler(score repository.Score) Manager {
	return Manager{scoreRepository: score}
}

func (m Manager) Handle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			m.processPost(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (m Manager) processPost(w http.ResponseWriter, r *http.Request) {
	score := server.Score{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		_, _ = w.Write([]byte("Errow while processing data"))
		return
	}

	scoreId, err := m.createScore(score)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("Error while processing data score"))
		return
	}

	_, _ = w.Write([]byte(scoreId))
}

func (m Manager) createScore(s server.Score) (string, error) {
	scoreEntity := repository.ScoreEntity{
		Score:     s.Score,
		RiderId:   s.RiderId,
		CreatedAt: time.Now(),
	}

	scoreId, err := m.scoreRepository.Save(scoreEntity)
	if err != nil {
		return "", err
	}

	return scoreId, nil
}
