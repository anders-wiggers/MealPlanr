package v1

import (
	controllerV1 "src/controllers/v1"

	"github.com/gin-gonic/gin"
)

//Initiate Routes for Version 1 of the API
func InitRoutes(g *gin.RouterGroup){
	g.GET("/",controllerV1.DefaultResponse)
	g.GET("/")	
}

