package mapper

import (
	"github.com/T-jegou/myTravelNotebook/pkg/entity"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
	"github.com/go-openapi/strfmt"
)

func MapTravel(e *entity.Travel) (*models.Travel, error) {
	r := &models.Travel{}
	r.ID = int64(e.ID)
	r.Name = &e.NameTravel
	r.Country = &e.Country
	r.Description = &e.Description
	r.CreationDate = strfmt.DateTime(e.CreatedAt)
	r.UpdatedAt = strfmt.DateTime(e.UpdatedAt)

	return r, nil
}
