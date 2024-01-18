package Controllers

import (
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllFoodItems godoc
// @Summary Get all food items or a specific item by EAN
// @Description Get all food items from the database or a specific item by ID or EAN
// @Accept json
// @Produce json
// @Param id path string false "Specifiy EAN to get a specific item"
// @Success 200 {object} []Model.Item
// @Router /api/getItems/{id} [get]
func GetAllFoodItems(c *gin.Context) {
	db := Config.DB

	ean := c.Query("ean")

	if ean != "" {
		// If it's not an ID, assume it's an EAN
		var items []Model.Item
		result := db.Where("\"item_ean\" = ?", ean).Find(&items)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, items)
		return
	}

	var items []Model.Item
	db.Find(&items)
	c.JSON(http.StatusOK, items)
}
