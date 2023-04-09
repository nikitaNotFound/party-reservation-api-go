package main

import (
	"net/http"

	_ "github.com/nikitaNotFound/party-reservation-api-go/users-service/cmd/docs"
	"github.com/nikitaNotFound/party-reservation-api-go/users-service/pkg/controllers/registration"
	"github.com/nikitaNotFound/party-reservation-api-go/users-service/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Party Reservation API
// @version         1.0
// @description     API for party tickets reservation
func main() {
	viper.SetConfigFile("../pkg/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbHost := viper.Get("DB_URL").(string)

	database := db.Init(dbHost)
	r := gin.Default()

	registration.RegisterRoutes(r, database)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		c.Abort()
	})

	r.Run(port)
}
