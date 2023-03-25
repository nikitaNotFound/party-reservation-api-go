package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-api/v2/cmd/docs"
	"go-api/v2/pkg/controllers/registration"
	"go-api/v2/pkg/db"
)

// @title           Party Reservation API
// @version         1.0
// @description     API for party tickets reservation
func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbHost := viper.Get("DB_URL").(string)

	database := db.Init(dbHost)
	r := gin.Default()

	registration.RegisterRoutes(r, database)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(port)
}
