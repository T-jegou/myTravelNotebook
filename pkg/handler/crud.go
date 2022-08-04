package handler

import (
	"github.com/AppDirect/argostatus/swagger_gen/restapi/operations/app"

	"github.com/go-openapi/runtime/middleware"
)

// CRUD is the CRUD interface
type CRUD interface {
	// apps
	GetApp(app.GetAppParams) middleware.Responder
	ListApps(app.ListAppsParams) middleware.Responder
}

// NewCRUD creates a new CRUD instance
func NewCRUD() CRUD {
	return &crud{}
}

type crud struct{}

func (c *crud) GetApp(params app.GetAppParams) middleware.Responder {
	return nil
}

func (c *crud) ListApps(params app.ListAppsParams) middleware.Responder {
	return nil
}
