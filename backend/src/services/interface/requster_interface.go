package services

// Interface for the Database requester
type DatabaseRequester interface {
	// Returns the type of database implemented
	GetDatabaseType() string

	//

}
