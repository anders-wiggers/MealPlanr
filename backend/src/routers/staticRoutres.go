package routers

//Static Routes all versions of the API should adhere to
type StaticRoutes struct {
	UserRoute         string `default:"/user"`
	RecipesRoute      string `deafult:"/user/recipes"`
	RecipeRoute       string `deafult:"/user/recipe"`
	PlannerRoute      string `deafult:"/planner"`
	ShoppingListRoute string `deafult:"/planner/shoppinglist"`
	LoginRoute        string `deafult:"/login"`
}

func NewStaticRoutes() *StaticRoutes {
	router := StaticRoutes{}

	router.UserRoute = "/user"
	router.RecipesRoute = "/user/recipes"
	router.RecipeRoute = "/user/recipe"
	router.PlannerRoute = "/planner"
	router.ShoppingListRoute = "/planner/shoppinglist"
	router.LoginRoute = "/login"

	return &router
}
