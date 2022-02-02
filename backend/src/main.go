package main

import (
	"fmt"
	"net/http"
	routersV1 "src/routers/v1"
	"src/services"

	"github.com/gin-gonic/gin"
)

//Gin Router
var router *gin.Engine

// Function responsiable for initiating the project
// Inint database
// Setup routing
func init() {
	services.GetDatabase()
	router = gin.New()
	version1 := router.Group("/v1")
	routersV1.InitRoutes(version1)
}

func main() {
	fmt.Println("Server Running on Port: ", 9090)
	http.ListenAndServe(":9090", router)
}
