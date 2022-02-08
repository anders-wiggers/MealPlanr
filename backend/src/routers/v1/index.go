package v1

import (
	controllerV1 "src/controllers/v1"
	"src/routers"

	"github.com/gin-gonic/gin"
)

//Initiate Routes for Version 1 of the API
func InitRoutes(g *gin.RouterGroup) {
	// Default test route
	g.GET("/", controllerV1.DefaultResponse)

	routes := routers.NewStaticRoutes()

	// User routes
	g.GET(routes.UserRoute, controllerV1.LookUpUser)
	g.POST(routes.UserRoute, controllerV1.DefaultResponse)  //TODO Implemnt
	g.PATCH(routes.UserRoute, controllerV1.DefaultResponse) //TODO Implement

	// User recipes routes
	g.GET(routes.RecipesRoute, controllerV1.DefaultResponse) //TODO Implement
	g.GET(routes.RecipeRoute, controllerV1.GetRecipe)
	g.POST(routes.RecipeRoute, controllerV1.UploadRecipes)
	g.DELETE(routes.RecipeRoute, controllerV1.DeleteRecipe)

	// Planner Routres
	g.GET(routes.PlannerRoute, controllerV1.DefaultResponse)  //TODO Implement
	g.POST(routes.PlannerRoute, controllerV1.DefaultResponse) //TODO Implement

	// ShoppingList routes
	g.GET(routes.ShoppingListRoute, controllerV1.DefaultResponse)   //TODO Implement
	g.PATCH(routes.ShoppingListRoute, controllerV1.DefaultResponse) //TODO Implement

	// Login Routes
	g.POST(routes.LoginRoute, controllerV1.DefaultResponse) //TODO Implement
}
