package Controllers

import (
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetAllFoodItems godoc
// @Summary Get all food items or a specific item by ID
// @Description Get all food items from the database or a specific item by ID
// @Tags get
// @Accept json
// @Produce json
// @Param id path int false "Item ID to get a specific item"
// @Success 200 {object} []Model.Item
// @Router /api/getItems/{id} [get]
func GetItems(c *gin.Context) {
	db := Config.DB

	if idParam, exists := c.Params.Get("id"); exists {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		var item Model.Item
		result := db.First(&item, id)
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
			return
		}

		c.JSON(http.StatusOK, item)
	} else {
		// If "id" parameter is not specified, get all food items
		var items []Model.Item
		db.Find(&items)

		c.JSON(http.StatusOK, items)
	}
}
