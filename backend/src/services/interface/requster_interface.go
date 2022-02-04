package services

import (
	"src/core/models"
)

// Interface for the Database requester
type DatabaseRequester interface {
	// Returns the type of database implemented
	GetDatabaseType() string

	// Return a user given the bearerToken
	GetUser(string) models.User

	// Updates a user given the bearerToken and stuct of info to update
	UpdateUser(string, models.User) string

	// Create a new recipe for a user with a bearerToken
	CreateRecipe(string, string, models.Recipe) string

	// Get a recipe from an ID
	GetRecipe(string) models.Recipe

	// Verify the validity of a bearerToken
	VerifyBearerToken(string) bool
}
