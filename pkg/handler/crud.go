package handler

import (
	"fmt"
	"reflect"

	"github.com/T-jegou/myTravelNotebook/pkg/entity"
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
	fmt.Println("Salutation")
	return nil
}

func (c *crud) DeleteTravel(params travel.DeleteTravelByIDParams) middleware.Responder {
	return nil
}

func (c *crud) UpdateTravel(params travel.UpdateTravelByIDParams) middleware.Responder {
	return nil
}

// travels
func (c *crud) GetTravels(params travel.GetTravelsParams) middleware.Responder {
	result := entity.GetDB().First(&entity.Travel{})
	fmt.Println("type du resultat renvoyé  : ", reflect.TypeOf(result))
	return nil
}

func (c *crud) AddTravel(params travel.AddTravelParams) middleware.Responder {
	println("hot  steuplait")

	return nil
}
