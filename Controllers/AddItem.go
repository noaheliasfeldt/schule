package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"image"
	"net/http"
)
import _ "image/png"
import _ "image/jpeg"

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func AddItem(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	src, srcErr := file.Open()
	if srcErr != nil {
		c.JSON(500, gin.H{"error1": srcErr.Error()})
		return
	}
	defer src.Close()

	imgDecode, _, error3 := image.Decode(src)
	if error3 != nil {
		c.JSON(400, gin.H{"error3": error3.Error()})
		return
	}

	bmp, error4 := gozxing.NewBinaryBitmapFromImage(imgDecode)
	if error4 != nil {
		c.JSON(400, gin.H{"error4": error4.Error()})
		return
	}

	ean13Reader := oned.NewEAN13Reader()
	result, _ := ean13Reader.Decode(bmp, nil)
	fmt.Println(ean13Reader)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
