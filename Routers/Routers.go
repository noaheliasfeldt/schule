package Routers

import (
	"feldtstudie-backend/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("addItem", Controllers.AddItem)
		api.POST("itemcategory", Controllers.AddItemCategory)
		api.GET("getItems", Controllers.GetItems)
		api.DELETE("deleteItems", Controllers.DeleteItem)
		api.GET("deleteItems", Controllers.ViewItemsExpiringSoon)

	}

	return r
}
