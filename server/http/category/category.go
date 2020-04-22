package category

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/BorsaTeam/jams-manager/server"
)

var categories = server.Categories{}

type Manager struct {
}

func NewCategoryHandler() Manager {
	return Manager{}
}

func (m Manager) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			processFindAll(w)
		case http.MethodPost:
			processPost(w, r)
		case http.MethodDelete:
			processDelete(r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func processPost(w http.ResponseWriter, r *http.Request) {
	uid, _ := uuid.NewRandom()
	category := server.Category{Id: uid.String()}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Error while processing data", http.StatusBadRequest)
	}

	categories = append(categories, category)
}

func processFindAll(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func processDelete(r *http.Request) {
	id := id(r.URL.Path)
	for i := range categories {
		if categories[i].Id == id {
			categories = append(categories[:i], categories[i+1:]...)
			break
		}
	}
}

func id(path string) string {
	ss := strings.Split(path, "/")
	if len(ss) == 3 {
		return ss[2]
	}
	return ""
}
