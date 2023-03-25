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

	r.POST("/registration", h.RegisterHandler)
}
