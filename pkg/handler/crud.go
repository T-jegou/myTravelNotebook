package handler

import (
	"github.com/T-jegou/myTravelNotebook/pkg/entity"
	"github.com/T-jegou/myTravelNotebook/pkg/mapper"
	"github.com/T-jegou/myTravelNotebook/pkg/util"
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
	GetTravels(travel.GetTravelsParams, *models.Principal) middleware.Responder

	// travel
	AddTravel(travel.AddTravelParams, *models.Principal) middleware.Responder
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

	if err := entity.GetDB().Unscoped().Delete(&entity.Travel{}, params.TravelID).Error; err != nil {
		return travel.NewDeleteTravelByIDDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	return travel.NewDeleteTravelByIDOK()
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
	t := &entity.Travel{}

	if params.Body != nil {
		t.NameTravel = util.SafeString(params.Body.Name)
		t.Description = util.SafeString(params.Body.Description)
		t.Country = util.SafeString(params.Body.Country)
	}

	tx := entity.GetDB().Begin()

	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return travel.NewAddTravelDefault(500).WithPayload(
			ErrorMessage("cannot create Travel. %s", err))
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return travel.NewAddTravelDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	resp := travel.NewAddTravelOK()

	return resp
}

// Health
func (c *crud) GetHealth(params health.GetHealthParams, principal *models.Principal) middleware.Responder {
	println("Get health")

	return nil
}
