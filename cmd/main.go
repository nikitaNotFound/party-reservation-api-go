package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-api/v2/pkg/controllers/registration"
	"go-api/v2/pkg/db"
)

func main() {
	viper.SetConfigFile("./pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbHost := viper.Get("DB_URL").(string)

	database := db.Init(dbHost)
	r := gin.Default()

	registration.RegisterRoutes(r, database)

	r.Run(port)
}
