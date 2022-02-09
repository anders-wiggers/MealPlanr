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
	standartRecipe.Date = "22-12-2022"
	standartRecipe.Author = "Anders Wiggers"
	standartRecipe.BasicInfo = models.BasicRecipeInfo{"10 mins", "25 mins", "Easy"}
	standartRecipe.CoverImage = "http://linktoimg.com/10001.png"

	var goose = models.IngredientsStuct{"Goose", "2000kg", true}
	var water = models.IngredientsStuct{"water", "200ml", false}
	var ingredientSlice = []models.IngredientsStuct{}
	ingredientSlice = append(ingredientSlice, goose)
	ingredientSlice = append(ingredientSlice, water)

	standartRecipe.Ingredients = ingredientSlice

	standartRecipe.Nutrition = models.NutritionStruct{"100", "12g", "4g", "100g", "0g", "3g", "20g", "0.5g"}

	standartRecipe.Method = append(standartRecipe.Method, "warm water")
	standartRecipe.Method = append(standartRecipe.Method, "put goose in water")
	standartRecipe.Method = append(standartRecipe.Method, "take goose out of water")

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
