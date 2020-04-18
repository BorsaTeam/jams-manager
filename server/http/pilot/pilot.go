package pilot

import (
	"encoding/json"
	"fmt"
	"github.com/BorsaTeam/jams-manager/server"
	"net/http"
	"strings"
)

var pilots = server.Pilots{}

type Manager struct {
}

func NewHandler() Manager {
	return Manager{}
}

func (m Manager) Handle() http.HandlerFunc  {


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
		}

	}
}

func processGet(w http.ResponseWriter, r *http.Request) {
	if id := id(r.URL.Path); id != "" {
		findOne(w, r)
		return
	}

	findAll(w)

}

func findAll(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pilots)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	id := id(r.URL.Path)
	var pilot *server.Pilot

	for i := range pilots {
		if pilots[i].Id == id {
			pilot = &pilots[i]
			break
		}
	}
	if pilot == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pilot)
}

func processPost(w http.ResponseWriter, r *http.Request) {
	pilot := server.Pilot{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&pilot)
	if err != nil {
		w.Write([]byte("Error while processing data"))
	}

	pilots = append(pilots, pilot)
}

func processPut(w http.ResponseWriter, r *http.Request) {
	pilot := server.Pilot{}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&pilot)
	if err != nil {
		w.Write([]byte("Error while processing data"))
	}

	id := id(r.URL.Path)
	fmt.Print()
	for i := range pilots {
		if pilots[i].Id == id {
			pilots[i] = pilot
		}
	}
}

func id(path string) string {
	if id := strings.Replace(path, "/pilots/", "", 1); id != "/pilots" {
		return id
	}
	return ""
}

func processDelete(r *http.Request) {
	id := id(r.URL.Path)

	for i := range pilots {
		if pilots[i].Id == id {
			pilots = append(pilots[:i], pilots[i+1:]...)
			break
		}
	}

}
