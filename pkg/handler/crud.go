package handler

import (
	"fmt"
	"time"

	"github.com/T-jegou/myTravelNotebook/swagger_gen/restapi/operations/travel"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type Travel struct {
	ID                int `gorm:"primaryKey;autoIncrement"`
	nameTravel        string
	descriptionTravel string
	country           string
	CreatedAt         time.Time `gorm:"autoCreateTime"`
}

func InitialMigration() {

	dsn := "host=localhost user=postgres password=postgrespw dbname=myTravel_notebook port=55000 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	} else {
		fmt.Println("SUCCESFULY CONNECTEC")
	}
	db.AutoMigrate(&Travel{})

	travel1 := Travel{
		nameTravel:        "Ptit voyage en pologne",
		descriptionTravel: "Pendant une semaine avec ma copine on va s'éclater à Cracovie",
		country:           "Pologne",
		CreatedAt:         time.Now(),
	}

	result := db.Create(&travel1)

	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println(travel1.ID)
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
	InitialMigration()
	return nil
}

func (c *crud) AddTravel(params travel.AddTravelParams) middleware.Responder {
	println("Bruhh")

	return nil
}
