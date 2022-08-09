package entity

import (
	"fmt"
	"sync"

	"github.com/T-jegou/myTravelNotebook/pkg/config"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	retry "github.com/avast/retry-go"
	// "github.com/jinzhu/gorm"
)

var (
	singletonDB   *gorm.DB
	singletonOnce sync.Once
)

// AutoMigrateTables stores the entity tables that we can auto migrate in gorm
var AutoMigrateTables = []interface{}{
	Travel{},
}

func connectDB() (db *gorm.DB, err error) {
	err = retry.Do(
		func() error {
			db, err = gorm.Open(config.Config.DBDriver, config.Config.DBConnectionStr)
			// db, err = gorm.Open("postgres", "postgres://user:password@host:5432/flagr?sslmode=disable")
			return err
		},
		retry.Attempts(config.Config.DBConnectionRetryAttempts),
		retry.Delay(config.Config.DBConnectionRetryDelay),
	)
	return db, err
}

// GetDB gets the db singleton
func GetDB() *gorm.DB {
	singletonOnce.Do(func() {
		db, err := connectDB()
		if err != nil {
			if config.Config.DBConnectionDebug {
				fmt.Print("Connexion string :", config.Config.DBConnectionStr)
				logrus.WithField("err", err).Fatal("failed to connect to db")
			} else {
				logrus.Fatal("failed to connect to db")
			}
		}
		db.Debug().AutoMigrate(AutoMigrateTables...)
		singletonDB = db
	})

	return singletonDB
}
