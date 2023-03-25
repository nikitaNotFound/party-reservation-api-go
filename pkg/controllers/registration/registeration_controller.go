package registration

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler{Db: db}

	routes := r.Group("/registration")

	routes.POST("/registration", h.RegisterHandler)
}
