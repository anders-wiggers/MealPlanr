package services

import (
	"src/services/db"
	services "src/services/interface"
)

/*
* 	Class responsible for creating the database
*	factory which creates the different instaces
*	of the database handlers
 */

// Strategy
type databaseHandler struct {
	DatabaseRequester services.DatabaseRequester
}

func DatabaseFactory() databaseHandler {
	return databaseHandler{
		DatabaseRequester: &db.DatabaseStub{},
	}
}
