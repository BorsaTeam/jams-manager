package database

import "database/sql"

type DbConnection interface {
	ConnectHandle() *sql.DB
	TestConnection()
}

type Migration interface {
	Apply()
}
