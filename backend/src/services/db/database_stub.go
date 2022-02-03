package db

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

func (databaseStub DatabaseStub) GetUser(string) struct{} {
	return struct{}{}
}

func (databaseStub DatabaseStub) UpdateUser(string, struct{}) string {
	return "succes"
}

func (databaseStub DatabaseStub) CreateRecipe(string, string, struct{}) string {
	return "succes"
}

func (databaseStub DatabaseStub) GetRecipe(string) struct{} {
	return struct{}{}
}
