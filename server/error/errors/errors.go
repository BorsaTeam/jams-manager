package errors

import "github.com/BorsaTeam/jams-manager/server"

var (
	Unknown = server.JamsError{
		Code:    "JAMS-001",
		Message: "There was an unknown error processing your request.",
	}
	PageRequestErr = server.JamsError{
		Code:    "JAMS-002",
		Message: `The query params "page" or "per_page" cannot be empty.`,
	}
)
