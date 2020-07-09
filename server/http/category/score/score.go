package score

import (
_ "encoding/json"
"net/http"
_ "net/http"
_ "strings"
"github.com/BorsaTeam/jams-manager/server"
)

var score = server.Score{}

type score struct {
}

func NewHandler() score {
	return score{}
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

