package models

/*
* 	Module for creating a Shopping List model
 */

type ShoppingList struct {
	List []Item `json:"list" binding:"required"`
}

func NewShoppingList() (shoppingList ShoppingList) {
	return shoppingList
}

// Helper structs
type Item struct {
	Name    string `json:"name" binding:"required"`
	Amount  string `json:"amount" binding:"required"`
	Checked bool   `json:"checked" binding:"required"`
}
