package v1

import (
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
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.HTML(http.StatusUnauthorized, "badaction.html", nil)
	}
}

func GetRecipe(c *gin.Context) {
	recipeId := c.Query("id")
	if recipeId == "" {
		c.HTML(http.StatusNotFound, "badaction.html", nil)
	} else {
		recipe := services.GetDatabaseInstance().DatabaseRequester.GetRecipe(recipeId)
		c.IndentedJSON(http.StatusOK, recipe)
	}
}

func DeleteRecipe(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "succes")
	//Check Do DB lookup for bearerToken to string
	bearerToken := c.GetHeader("bearerToken")

	services.GetDatabaseInstance().DatabaseRequester.DeleteRecipe(bearerToken, "ID") //TODO FIX ID FROM CONTENT
}

func UploadRecipes(c *gin.Context) {

}
