package main

import (
	"feldtstudie-backend/Config"
	"feldtstudie-backend/Model"
	"feldtstudie-backend/Routers"
	docs "feldtstudie-backend/docs"
	"fmt"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("an upsi happend: ", err)
		return
	}
	err := Config.DB.AutoMigrate(
		&Model.Item{},
		&Model.ItemCategory{},
	)
	if err != nil {
		return
	}
	docs.SwaggerInfo.BasePath = "/api/"
	r := Routers.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// running
	r.Run()
}
