package services

import (
	"src/services/db"
	services "src/services/interface"
)

// Strategy
type databaseHandler struct {
	DatabaseRequester services.DatabaseRequester
}

func DatabaseFactory() databaseHandler {
	return databaseHandler{
		DatabaseRequester: &db.DatabaseStub{},
	}
}
