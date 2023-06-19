package events

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitaNotFound/party-reservation-api-go/users-service/pkg/controllers"
	"gorm.io/gorm"
)

type handler struct {
	*controllers.UserServiceHandler
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler{
		&controllers.UserServiceHandler{
			Db: db,
		},
	}

	r.POST("/registration", h.RegisterHandler)
}
