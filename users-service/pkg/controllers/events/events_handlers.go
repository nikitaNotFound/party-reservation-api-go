package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitaNotFound/party-reservation-api-go/users-service/pkg/models"
)

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// RegisterHandler godoc
// @Summary Registration to the party
// @Schemes
// @Description process incoming data to register use to the party. Sends email or sms with reservation code
// @Tags registration
// @Accept json
// @Produce json
// @Router /registration [post]
func (h handler) RegisterHandler(c *gin.Context) {
	request := RegisterRequest{}

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	registration := models.Registration{
		ID:      0,
		UserID:  0,
		PartyID: 0,
	}

	if result := h.Db.Create(&registration); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &registration)
}
