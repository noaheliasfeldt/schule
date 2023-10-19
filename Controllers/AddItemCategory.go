package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddItemCategory(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "lol",
	})
}
