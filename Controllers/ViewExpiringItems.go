package Controllers

import (
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary view items expiring soon
// @Description View items that expire in the next 3 days or earlier
// @Tags view
// @Accept json
// @Produce json
// @Success 200
// @Router /api/viewitems [get]
func ViewItemsExpiringSoon(c *gin.Context) {
	db := Config.DB

	// Berechne das Datum, das 3 Tage in der Zukunft liegt
	threeDaysFromNow := time.Now().Add(3 * 24 * time.Hour)

	// Abfrage für Artikel, die in den nächsten 3 Tagen oder früher ablaufen
	var expiringItems []Model.Item
	result := db.Where("ItemBBD <= ?", threeDaysFromNow.Unix()).Find(&expiringItems)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": expiringItems})
}
