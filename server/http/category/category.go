package category

import (
	"encoding/json"
	"github.com/BorsaTeam/jams-manager/server"
	"net/http"
	"strings"
)

var categories = server.Categories{}

func Handle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			processGet(w)
		case http.MethodPost:
			processPost(w, r)
		case http.MethodDelete:
			processDelete(r)
		default:
			http.NotFound(w, r)
		}
	}
}

func processPost(w http.ResponseWriter, r *http.Request) {

	category := server.Category{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.Write([]byte("Error while processing data"))
	}

	categories = append(categories, category)
}

func processGet(w http.ResponseWriter) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func serviceFromPath(path string) string {
	return strings.Replace(path, "/categories/", "", 1)
}

func processDelete(r *http.Request) {
	id := serviceFromPath(r.URL.Path)

	for i := range categories {
		if categories[i].Id == id {
			categories = append(categories[:i], categories[i+1:]...)
			break
		}
	}
}

