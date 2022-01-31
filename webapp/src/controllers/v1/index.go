package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultResponse(c *gin.Context){
	c.JSON(http.StatusOK, "hihi")
}