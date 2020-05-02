package rider

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/BorsaTeam/jams-manager/server"
)

var riders = server.Riders{}

type Manager struct {
}

func NewHandler() Manager {
	return Manager{}
}

func (m Manager) Handle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			processGet(w, r)
		case http.MethodPost:
			processPost(w, r)
		case http.MethodDelete:
			processDelete(r)
		case http.MethodPut:
			processPut(w, r)
		default:
			http.Error(w,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
		}
	}
}

func processGet(w http.ResponseWriter, r *http.Request) {
	if id := id(r.URL.Path); id != "" {
		rider := findOne(id)
		if rider.Id == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rider)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(findAll())
}

func findOne(id string) server.Rider {

	for i := range riders {
		if riders[i].Id == id {
			return riders[i]
		}
	}
	return server.Rider{}
}

func findAll() server.Riders {
	return riders
}

func processPost(w http.ResponseWriter, r *http.Request) {
	rider := server.Rider{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rider)
	if err != nil {
		w.Write([]byte("Error while processing data"))
	}

	riders = append(riders, rider)
}

func processPut(w http.ResponseWriter, r *http.Request) {
	rider := server.Rider{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&rider)
	if err != nil {
		w.Write([]byte("Error while processing data"))
	}

	id := id(r.URL.Path)
	for i := range riders {
		if riders[i].Id == id {
			riders[i] = rider
		}
	}
}

func id(path string) string {
	p := strings.Split(path, "/")
	if len(p) > 1 {
		return p[2]
	}
	return ""
}

func processDelete(r *http.Request) {
	id := id(r.URL.Path)

	for i := range riders {
		if riders[i].Id == id {
			riders = append(riders[:i], riders[i+1:]...)
			break
		}
	}
}
