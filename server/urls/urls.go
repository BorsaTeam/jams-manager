package urls

import (
	"net/http"
	"strconv"

	"github.com/BorsaTeam/jams-manager/server"
	"github.com/BorsaTeam/jams-manager/server/error/errors"
)

func Page(r *http.Request) (server.PageRequest, error) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	perPageStr := query.Get("per_page")

	if pageStr == "" || perPageStr == "" {
		return server.PageRequest{}, errors.PageRequestErr
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return server.PageRequest{}, errors.PageRequestErr
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		return server.PageRequest{}, errors.PageRequestErr
	}

	pageReq := server.PageRequest{
		Page:    page,
		PerPage: perPage,
	}

	return pageReq, nil
}
