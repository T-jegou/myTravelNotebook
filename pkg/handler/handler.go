package handler

import (
	"github.com/AppDirect/argostatus/swagger_gen/models"
	"github.com/AppDirect/argostatus/swagger_gen/restapi/operations"
	"github.com/AppDirect/argostatus/swagger_gen/restapi/operations/app"
	"github.com/AppDirect/argostatus/swagger_gen/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

// Setup initialize all the handler functions
func Setup(api *operations.ArgostatusAPI) {
	setupHealth(api)
	setupCRUD(api)
}

func setupCRUD(api *operations.ArgostatusAPI) {
	c := NewCRUD()
	// flags
	api.AppGetAppHandler = app.GetAppHandlerFunc(c.GetApp)
	api.AppListAppsHandler = app.ListAppsHandlerFunc(c.ListApps)
}

func setupHealth(api *operations.ArgostatusAPI) {
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(
		func(health.GetHealthParams) middleware.Responder {
			return health.NewGetHealthOK().WithPayload(&models.Health{Status: "OK"})
		},
	)
}
