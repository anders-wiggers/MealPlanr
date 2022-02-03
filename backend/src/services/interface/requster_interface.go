package services

// Interface for the Database requester
type DatabaseRequester interface {
	// Returns the type of database implemented
	GetDatabaseType() string

	// Return a user given the bearerToken
	GetUser(string) struct{}

	// Updates a user given the bearerToken and stuct of info to update
	UpdateUser(string, struct{}) string

	// Create a new recipe for a user with a bearerToken
	CreateRecipe(string, string, struct{}) string

	// Get a recipe from an ID
	GetRecipe(string) struct{}
}
