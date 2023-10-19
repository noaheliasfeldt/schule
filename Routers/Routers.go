package Routers

import (
	"feldtstudie-backend/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("item", Controllers.AddItem)
		api.POST("itemcategory", Controllers.AddItemCategory)
	}

	return r
}
