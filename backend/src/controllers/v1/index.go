package v1

import (
	"fmt"
	"net/http"
	"src/services"

	"github.com/gin-gonic/gin"
)

func DefaultResponse(c *gin.Context){
	c.JSON(http.StatusOK, "hihi")
}

func LookUpUser(c *gin.Context){
	//Check Do DB lookup for bearerToken to string
	// TODO: IMPORTANT FIX DATABASE IMPORT 
	test := services.DatabaseFactory()	
	bearerToken := c.GetHeader("bearerToken")
	if ValidateBearerToken(bearerToken) {
		user := test.DatabaseRequester.GetUser(bearerToken)
		fmt.Println(bearerToken)
		c.IndentedJSON(http.StatusOK,user)
	} else {
		c.HTML(http.StatusUnauthorized,"badaction.html",nil)
	}
	
}