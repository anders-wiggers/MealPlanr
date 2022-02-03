package models

/*
* 	Module for creating a Week Planner
 */

type Planner struct {
	week string
	year string
	days []WeekDay
}

func NewPlanner() (planner Planner) {
	return planner
}

// Helper Structs
type WeekDay struct {
	date      string
	breakfast Meal
	launch    Meal
	dinner    Meal
}

type Meal struct {
	status  string
	meal_id int
}
