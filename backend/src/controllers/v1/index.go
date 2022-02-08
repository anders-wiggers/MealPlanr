package v1

import (
	"fmt"
	"net/http"
	"src/services"

	"github.com/gin-gonic/gin"
)

func DefaultResponse(c *gin.Context) {
	c.JSON(http.StatusOK, "hihi")
}

func LookUpUser(c *gin.Context) {
	//Check Do DB lookup for bearerToken to string
	bearerToken := c.GetHeader("bearerToken")
	if ValidateBearerToken(bearerToken) {
		user := services.GetDatabaseInstance().DatabaseRequester.GetUser(bearerToken)
		fmt.Println(bearerToken)
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.HTML(http.StatusUnauthorized, "badaction.html", nil)
	}
}

func GetRecipe(c *gin.Context) {

}

func DeleteRecipe(c *gin.Context) {

}

func UploadRecipes(c *gin.Context) {

}
