package models

/*
* 	Module for creating a Week Planner
 */

type Planner struct {
	Week string    `json:"week" binding:"required"`
	Year string    `json:"year" binding:"required"`
	Days []WeekDay `json:"days" binding:"required"`
}

func NewPlanner() (planner Planner) {
	return planner
}

// Helper Structs
type WeekDay struct {
	Date      string `json:"date" binding:"required"`
	Breakfast Meal   `json:"breakfast" binding:"required"`
	Launch    Meal   `json:"launch" binding:"required"`
	Dinner    Meal   `json:"dinner" binding:"required"`
}

type Meal struct {
	Status  string `json:"status" binding:"required"`
	Meal_id int    `json:"mealId" binding:"required"`
}
