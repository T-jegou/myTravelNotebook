package handler

import (
	"fmt"

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
	fmt.Println(params.HTTPRequest.Body)
	return nil
}

func (c *crud) AddTravel(params travel.AddTravelParams) middleware.Responder {
	println("Bruhh")

	return nil
}

/* func (c *crud) CreateTag(params tag.CreateTagParams) middleware.Responder {
	s := &entity.Flag{}
	s.ID = uint(params.FlagID)
	t := &entity.Tag{}
	t.Value = util.SafeString(params.Body.Value)
	if ok, reason := util.IsSafeValue(t.Value); !ok {
		return tag.NewCreateTagDefault(400).WithPayload(ErrorMessage("%s", reason))
	}

	getDB().Where("value = ?", util.SafeString(params.Body.Value)).Find(t) // Find the existing tag to associate if it exists
	if err := getDB().Model(s).Association("Tags").Append(t).Error; err != nil {
		return tag.NewCreateTagDefault(500).WithPayload(ErrorMessage("%s", err))
	}

	resp := tag.NewCreateTagOK()
	resp.SetPayload(e2r.MapTag(t))

	entity.SaveFlagSnapshot(getDB(), util.SafeUint(params.FlagID), getSubjectFromRequest(params.HTTPRequest))
	return resp
}
*/
