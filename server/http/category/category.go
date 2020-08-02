package category

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/database/repository"
)

type Manager struct {
	repo repository.Category
}

var test = ""

func NewHandler(repo repository.Category) Manager {
	return Manager{repo: repo}
}

func (m Manager) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			m.processFindAll(w)
		case http.MethodPost:
			m.processPost(w, r)
		case http.MethodDelete:
			m.processDelete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (m Manager) processPost(w http.ResponseWriter, r *http.Request) {
	category := server.Category{}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Error while processing data", http.StatusBadRequest)
	}

	categoryId, err := m.createCategory(category)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("Error while processing data CATEGORY"))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	_, _ = w.Write([]byte(categoryId))
}

func (m Manager) createCategory(category server.Category) (server.CategoryId, error) {
	entity := repository.CategoryEntity{
		Name:     category.Name,
		CreateAt: time.Now(),
	}

	categoryId, err := m.repo.Save(entity)
	if err != nil {
		return "", err
	}

	return server.CategoryId(categoryId), nil
}

func (m Manager) processFindAll(w http.ResponseWriter) {
	categories, err := m.findAll()
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("Error while processing data CATEGORY"))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(categories)
}

func (m Manager) findAll() (server.Categories, error) {
	cc, err := m.repo.FindAll()
	if err != nil {
		return nil, err
	}

	categories := make(server.Categories, len(cc))
	for i := range cc {
		categories[i].Id = server.CategoryId(cc[i].Id)
		categories[i].Name = cc[i].Name
		categories[i].CreatedAt = cc[i].CreateAt
		categories[i].UpdatedAt = cc[i].UpdateAt
	}

	return categories, nil
}

func (m Manager) processDelete(w http.ResponseWriter, r *http.Request) {
	id := id(r.URL.Path)
	if err := m.delete(server.CategoryId(id)); err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("Error while processing data CATEGORY"))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
}

func (m Manager) delete(categoryId server.CategoryId) error {
	if err := m.repo.Delete(repository.CategoryId(categoryId)); err != nil {
		return err
	}

	return nil
}

func id(path string) string {
	ss := strings.Split(path, "/")
	if len(ss) == 3 {
		return ss[2]
	}
	return ""
}
