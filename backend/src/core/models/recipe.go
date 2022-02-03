package models

/*
* 	Module for creating a Recipe
 */

type Recipe struct {
	title       string
	uid         string
	date        string
	author      string
	basicInfo   basicRecipeInfo
	coverImage  string
	ingredients []ingredientsStuct
	nutrition   nutritionStruct
	Method      []string
}

func NewRecipe() (recipe Recipe) {
	return recipe
}

// Helper data structures
type basicRecipeInfo struct {
	prebTime   string
	cookTime   string
	difficulty string
}

type ingredientsStuct struct {
	name         string
	amount       string
	purchaseable bool
}

type nutritionStruct struct {
	kcal      string
	fat       string
	saturates string
	carbs     string
	sugars    string
	fibre     string
	protein   string
	salt      string
}
