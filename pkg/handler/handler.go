package handler

import (
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations"
	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/travel"
)

func Setup(api *operations.MyTravelNotebookAPI) {
	setupCRUD(api)
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
}
