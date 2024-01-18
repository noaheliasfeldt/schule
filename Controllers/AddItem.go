package Controllers

import (
	"encoding/json"
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)
import _ "image/png"
import _ "image/jpeg"

// @BasePath /api/v1

// PingExample godoc
// @Summary add item
// @Description Add an item by barcode number
// @Accept json
// @Produce json
// @Param ean query string true "EAN (barcode) number"
// @Param mhd query string true "MHD (expiration date)"
// @Param count query string true "Count (amount of item to add)"
// @Success 200
// @Router /api/additem [post]
func AddItem(c *gin.Context) {
	ean := c.Query("ean")
	mhd := c.Query("mhd")
	count := c.Query("count")

	apiURL := fmt.Sprintf("https://world.openfoodfacts.org/api/v2/product/%s.json", ean)

	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product information"})
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	var productInfo map[string]interface{}
	err = json.Unmarshal(body, &productInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response"})
		return
	}

	product, ok := productInfo["product"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product information is missing or not in the expected format"})
		return
	}

	productName, ok := product["product_name"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product name is missing or not a string"})
		return
	}

	mhdInt, err := strconv.ParseInt(mhd, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MHD format"})
		return
	}

	eanInt, err := strconv.ParseInt(ean, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MHD format"})
		return
	}

	db := Config.DB

	countUint64, err := strconv.ParseUint(count, 10, 64)
	if err != nil {
		fmt.Println("Fehler beim Konvertieren des Strings zu uint64:", err)
		return
	}

	var existingItem Model.Item
	result := db.Where("\"item_bbd\" = ? AND \"item_ean\" = ?", mhdInt, ean).First(&existingItem)
	if result.Error == gorm.ErrRecordNotFound {

		newItem := Model.Item{
			ItemName:  productName,
			ItemBBD:   mhdInt,
			ItemEAN:   eanInt, // Hinzufügen der EAN zum Vergleich
			ItemCount: uint(countUint64),
		}
		db.Create(&newItem)
	} else if result.Error == nil {
		existingItem.ItemCount += uint(countUint64) // Increase ItemCount by the specified count
		db.Save(&existingItem)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Item added successfully", "foodInfo": productInfo})
}
