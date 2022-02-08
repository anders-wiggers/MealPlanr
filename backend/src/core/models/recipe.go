package models

/*
* 	Module for creating a Recipe
 */

type Recipe struct {
	Title       string             `json:"title" binding:"required"`
	Uid         string             `json:"uid" binding:"required"`
	Date        string             `json:"date" binding:"required"`
	Author      string             `json:"author" binding:"required"`
	BasicInfo   basicRecipeInfo    `json:"basic_info" binding:"required"`
	CoverImage  string             `json:"cover_image" binding:"required"`
	Ingredients []ingredientsStuct `json:"ingredients" binding:"required"`
	Nutrition   nutritionStruct    `json:"nutritions" binding:"required"`
	Method      []string           `json:"method" binding:"required"`
}

func NewRecipe() (recipe Recipe) {
	return recipe
}

// Helper data structures
type basicRecipeInfo struct {
	PrebTime   string
	CookTime   string
	Difficulty string
}

type ingredientsStuct struct {
	Name         string
	Amount       string
	Purchaseable bool
}

type nutritionStruct struct {
	Kcal      string
	Fat       string
	Saturates string
	Carbs     string
	Sugars    string
	Fibre     string
	Protein   string
	Salt      string
}
