package handler

import (
	errors "github.com/go-openapi/errors"

	"github.com/T-jegou/myTravelNotebook/swagger_gen/models"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/authentication"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/health"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/login"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/travel"
)

func Setup(api *operations.MyTravelNotebookAPI) {
	setupCRUD(api)
	setupAuthent(api)
}

func setupAuthent(api *operations.MyTravelNotebookAPI) {
	api.OauthSecurityAuth = func(token string, scopes []string) (*models.Principal, error) {
		ok, err := authenticated(token)
		if err != nil {
			return nil, errors.New(401, "error authenticate")
		}
		if !ok {
			return nil, errors.New(401, "invalid token")
		}
		prin := models.Principal(token)
		return &prin, nil

	}
}

func setupCRUD(api *operations.MyTravelNotebookAPI) {
	c := NewCRUD()

	// Travels
	api.TravelGetTravelsHandler = travel.GetTravelsHandlerFunc(c.GetTravels)
	api.TravelAddTravelHandler = travel.AddTravelHandlerFunc(c.AddTravel)

	// Travel
	api.TravelGetTravelByIDHandler = travel.GetTravelByIDHandlerFunc(c.GetTravel)
	api.TravelDeleteTravelByIDHandler = travel.DeleteTravelByIDHandlerFunc(c.DeleteTravel)
	api.TravelUpdateTravelByIDHandler = travel.UpdateTravelByIDHandlerFunc(c.UpdateTravel)

	// Health
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(c.GetHealth)

	// Login
	api.LoginGetLoginHandler = login.GetLoginHandlerFunc(c.GetLogin)

	// Authentication
	api.AuthenticationGetAuthCallbackHandler = authentication.GetAuthCallbackHandlerFunc(c.GetAuthCallback)

}
