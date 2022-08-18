package handler

import (
	"fmt"

	"github.com/T-jegou/myTravelNotebook/pkg/entity"
	"github.com/T-jegou/myTravelNotebook/pkg/mapper"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/authentication"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/health"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/login"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/travel"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type CRUD interface {
	// Login
	GetLogin(login.GetLoginParams) middleware.Responder

	// Authentication
	GetAuthCallback(authentication.GetAuthCallbackParams) middleware.Responder

	// travels
	AddTravel(travel.AddTravelParams, *models.Principal) middleware.Responder
	GetTravels(travel.GetTravelsParams, *models.Principal) middleware.Responder

	// travel
	GetTravel(travel.GetTravelByIDParams, *models.Principal) middleware.Responder
	DeleteTravel(travel.DeleteTravelByIDParams, *models.Principal) middleware.Responder
	UpdateTravel(travel.UpdateTravelByIDParams, *models.Principal) middleware.Responder

	// Health
	GetHealth(health.GetHealthParams, *models.Principal) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	return &crud{}
}

type crud struct{}

// login
func (c *crud) GetLogin(params login.GetLoginParams) middleware.Responder {
	return Login(params.HTTPRequest)
}

// authentication
func (c *crud) GetAuthCallback(params authentication.GetAuthCallbackParams) middleware.Responder {
	token, err := Callback(params.HTTPRequest)
	if err != nil {
		return middleware.NotImplemented("operation .GetAuthCallback error")
	}
	logrus.Println("Token", token)
	return authentication.NewGetAuthCallbackDefault(500).WithPayload(
		ErrorMessage("Token %v", token))
}

// travel
func (c *crud) GetTravel(params travel.GetTravelByIDParams, principal *models.Principal) middleware.Responder {
	t := &entity.Travel{}
	result := entity.GetDB().First(t, params.TravelID)

	if result.RecordNotFound() {
		return travel.NewGetTravelByIDDefault(404).WithPayload(
			ErrorMessage("unable to find travel %v in the database", params.TravelID))
	}

	if err := result.Error; err != nil {
		return travel.NewGetTravelByIDDefault(500).WithPayload(
			ErrorMessage("Unknown error occurred while retrieve travel %v: %s", params.TravelID, err))
	}

	resp := travel.NewGetTravelByIDOK()
	payload, err := mapper.MapTravel(t)
	if err != nil {
		return travel.NewGetTravelByIDDefault(500).WithPayload(
			ErrorMessage("Cannot map travel %v. %s", params.TravelID, err))
	}
	resp.SetPayload(payload)
	return resp

}

func (c *crud) DeleteTravel(params travel.DeleteTravelByIDParams, principal *models.Principal) middleware.Responder {
	return nil
}

func (c *crud) UpdateTravel(params travel.UpdateTravelByIDParams, principal *models.Principal) middleware.Responder {
	return nil
}

// travels
func (c *crud) GetTravels(params travel.GetTravelsParams, principal *models.Principal) middleware.Responder {
	return nil
}

func (c *crud) AddTravel(params travel.AddTravelParams, principal *models.Principal) middleware.Responder {
	println("Add travel")
	dbInstance := entity.GetDB()
	travel := entity.Travel{NameTravel: "Escapade solitaire A Cracovie", Country: "Pologne", Description: "Blah blah blah"}

	result := dbInstance.Create(&travel)
	fmt.Print(result.Error, travel.ID)

	return nil
}

// Health
func (c *crud) GetHealth(params health.GetHealthParams, principal *models.Principal) middleware.Responder {
	println("Get health")

	return nil
}
