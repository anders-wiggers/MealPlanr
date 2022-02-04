package db

import (
	"src/core/models"
)

/*
Stub implementation of the database used for testing
The stub implementes static lookup that makes it possible
to create static test.
*/

// Default Struct
type DatabaseStub struct{}

func (databaseStub DatabaseStub) GetDatabaseType() string {
	return "stub"
}

func (databaseStub DatabaseStub) GetUser(string) models.User {
	return models.NewUser("test", "11", "https:33")
}

func (databaseStub DatabaseStub) UpdateUser(string, models.User) string {
	return "succes"
}

func (databaseStub DatabaseStub) CreateRecipe(string, string, models.Recipe) string {
	return "succes"
}

func (databaseStub DatabaseStub) GetRecipe(string) models.Recipe {
	return models.NewRecipe()
}

func (databaseStub DatabaseStub) VerifyBearerToken(string) bool {
	return true
}