package handler

import (
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
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	return &crud{}
}

type crud struct{}

// travel
func (c *crud) GetTravel(params travel.GetTravelByIDParams) middleware.Responder {
	return nil
}

func (c *crud) DeleteTravel(params travel.DeleteTravelByIDParams) middleware.Responder {
	return nil
}

func (c *crud) UpdateTravel(travel.UpdateTravelByIDParams) middleware.Responder {
	return nil
}

// travels
func (c *crud) GetTravels(params travel.GetTravelsParams) middleware.Responder {
	return nil
}

func (c *crud) AddTravel(travel.AddTravelParams) middleware.Responder {
	return nil
}
