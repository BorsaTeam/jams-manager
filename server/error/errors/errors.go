package errors

import "github.com/BorsaTeam/jams-manager/server"

var Unknown = server.JamsError{
	Code:    "JAMS-001",
	Message: "There was an unknown error processing your request.",
}
