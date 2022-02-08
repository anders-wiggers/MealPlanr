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

func (databaseStub DatabaseStub) GetRecipe(id string) models.Recipe {
	var standartRecipe = models.NewRecipe()
	standartRecipe.Title = "Boiled Goose"
	standartRecipe.Uid = id

	return standartRecipe
}

func (databaseStub DatabaseStub) VerifyBearerToken(bearerToken string) bool {
	if bearerToken == "5230-SF20b-&21c1-%8vs1x41sd" {
		return true
	} else {
		return false
	}
}

func (databaseStub DatabaseStub) DeleteRecipe(string, string) string {
	return "succes"
}
