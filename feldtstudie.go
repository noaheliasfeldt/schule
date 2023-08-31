package main

import (
	"feldtstudie-backend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err.Error() != "" {
		fmt.Println(err.Error())
	}

	r := gin.Default()
	r.GET("/addItem", routes.AddItem)
	r.Run()
}
