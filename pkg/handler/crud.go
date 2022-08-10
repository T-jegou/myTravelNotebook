package handler

import (
	"fmt"

	"github.com/T-jegou/myTravelNotebook/pkg/entity"
	"github.com/T-jegou/myTravelNotebook/pkg/mapper"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/health"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/travel"
	"github.com/go-openapi/runtime/middleware"
)

type CRUD interface {
	// travels
	AddTravel(travel.AddTravelParams) middleware.Responder
	GetTravels(travel.GetTravelsParams) middleware.Responder

	// travel
	GetTravel(travel.GetTravelByIDParams) middleware.Responder
	DeleteTravel(travel.DeleteTravelByIDParams) middleware.Responder
	UpdateTravel(travel.UpdateTravelByIDParams) middleware.Responder

	// Health
	GetHealth(health.GetHealthParams) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	return &crud{}
}

type crud struct{}

// travel
func (c *crud) GetTravel(params travel.GetTravelByIDParams) middleware.Responder {
	t := &entity.Travel{}
	result := entity.GetDB().First(t, params.TravelID)

	if result.RecordNotFound() {
		return travel.NewGetTravelByIDDefault(404).WithPayload(
			ErrorMessage("unable to find flag %v in the database", params.TravelID))
	}

	if err := result.Error; err != nil {
		return travel.NewGetTravelByIDDefault(500).WithPayload(
			ErrorMessage("Unknown error occurred while retrieve travel %v: %s", params.TravelID, err))
	}

	resp := travel.NewGetTravelByIDOK()
	payload, err := mapper.MapTravel(t)
	if err != nil {
		return travel.NewGetTravelByIDDefault(500).WithPayload(
			ErrorMessage("cannot map travel %v. %s", params.TravelID, err))
	}
	resp.SetPayload(payload)
	return resp

}

func (c *crud) DeleteTravel(params travel.DeleteTravelByIDParams) middleware.Responder {
	return nil
}

func (c *crud) UpdateTravel(params travel.UpdateTravelByIDParams) middleware.Responder {
	return nil
}

// travels
func (c *crud) GetTravels(params travel.GetTravelsParams) middleware.Responder {
	dbInstance := entity.GetDB()
	travel := entity.Travel{NameTravel: "Quelques jours en normandie", Country: "France", Description: "Blah blah blah"}

	result := dbInstance.Create(&travel)
	fmt.Print(result.Error, travel.ID)
	return nil
}

func (c *crud) AddTravel(params travel.AddTravelParams) middleware.Responder {
	println("Add travel")

	return nil
}

// Health
func (c *crud) GetHealth(params health.GetHealthParams) middleware.Responder {
	println("Get health")

	return nil
}
