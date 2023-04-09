package registration

import (
	"github.com/gin-gonic/gin"
	"go-api/v2/pkg/models"
	"net/http"
)

type RegisterRequest struct {
	FirstName        string                  `json:"firstName"`
	LastName         string                  `json:"lastName"`
	Age              int                     `json:"age"`
	RegistrationType models.RegistrationType `json:"registrationType"`
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
	body := RegisterRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	registration := models.Registration{
		FirstName:        body.FirstName,
		LastName:         body.LastName,
		Age:              body.Age,
		RegistrationType: body.RegistrationType,
	}

	if result := h.Db.Create(&registration); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &registration)
}
