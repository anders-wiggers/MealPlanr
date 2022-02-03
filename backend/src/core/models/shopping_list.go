package models

/*
* 	Module for creating a Shopping List model
 */

type ShoppingList struct {
	list []item
}

func NewShoppingList() (shoppingList ShoppingList) {
	return shoppingList
}

// Helper structs
type item struct {
	name    string
	amount  string
	checked bool
}
