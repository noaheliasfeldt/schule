package Controllers

import (
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// @Summary delete item
// @Description Delete an item by EAD, MHD, or ID
// @Accept json
// @Produce json
// @Param ead query string false "EAD (barcode) number"
// @Param mhd query string false "MHD (expiration date)"
// @Param id query uint false "Item ID"
// @Success 200
// @Router /api/deleteitem [delete]
func DeleteItem(c *gin.Context) {
	// EAD, MHD und ID aus den Abfrageparametern abrufen
	ead := c.Query("ead")
	mhd := c.Query("mhd")
	idStr := c.Query("id")

	db := Config.DB

	// Überprüfen, ob eine der Abfrageparameter vorhanden ist
	if ead == "" && mhd == "" && idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one of ead, mhd, or id must be provided"})
		return
	}

	// Entweder EAD, MHD oder ID sollte für die Abfrage verwendet werden
	var result *gorm.DB
	if ead != "" {
		result = db.Where("\"item_ean\" = ?", ead).Delete(&Model.Item{})
	} else if mhd != "" {
		result = db.Where("\"item_bbd\" = ?", mhd).Delete(&Model.Item{})
	} else {
		// Wenn ID vorhanden ist, versuche es zu konvertieren und das Item zu löschen
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		result = db.Where("ID = ?", id).Delete(&Model.Item{})
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item from the database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Item deleted successfully"})
}
