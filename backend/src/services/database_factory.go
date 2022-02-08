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
type DatabaseHandler struct {
	DatabaseRequester services.DatabaseRequester
}

func DatabaseFactory() *DatabaseHandler {
	return &DatabaseHandler{
		DatabaseRequester: &db.DatabaseStub{},
	}
}

var database *DatabaseHandler

// Singleton pattern to return a database instance
func GetDatabaseInstance() *DatabaseHandler {
	if database == nil {
		database = DatabaseFactory()
	}
	return database
}
