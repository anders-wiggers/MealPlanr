package models

/*
* 	Module for creating a Recipe
 */

type Recipe struct {
	Title       string             `json:"title" binding:"required"`
	Uid         string             `json:"uid" binding:"required"`
	Date        string             `json:"date" binding:"required"`
	Author      string             `json:"author" binding:"required"`
	BasicInfo   BasicRecipeInfo    `json:"basic_info" binding:"required"`
	CoverImage  string             `json:"cover_image" binding:"required"`
	Ingredients []IngredientsStuct `json:"ingredients" binding:"required"`
	Nutrition   NutritionStruct    `json:"nutritions" binding:"required"`
	Method      []string           `json:"method" binding:"required"`
}

func NewRecipe() (recipe Recipe) {
	return recipe
}

// Helper data structures
type BasicRecipeInfo struct {
	PrebTime   string `json:"prebTime" binding:"required"`
	CookTime   string `json:"cook" binding:"required"`
	Difficulty string `json:"difficulty" binding:"required"`
}

type IngredientsStuct struct {
	Name         string `json:"name" binding:"required"`
	Amount       string `json:"amount" binding:"required"`
	Purchaseable bool   `json:"purchasable" binding:"required"`
}

type NutritionStruct struct {
	Kcal      string `json:"kcal" binding:"required"`
	Fat       string `json:"fat" binding:"required"`
	Saturates string `json:"saturates" binding:"required"`
	Carbs     string `json:"carbs" binding:"required"`
	Sugars    string `json:"sugars" binding:"required"`
	Fibre     string `json:"fibre" binding:"required"`
	Protein   string `json:"protein" binding:"required"`
	Salt      string `json:"salt" binding:"required"`
}
