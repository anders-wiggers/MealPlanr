package db

type DatabaseStub struct{}

func (databaseStub DatabaseStub) GetDatabaseType() string {
	return "stub"
}
