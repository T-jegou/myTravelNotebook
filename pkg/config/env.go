package config

import "time"

/*
Anonymous structure ; check google
Ici on délcare les variables de configuration, soit on utilise des variables d'environnement qui
doivent être défini auparavant , soit on affecte un choix par défaut */
var Config = struct {

	// Database stuff
	DBDriver          string `env:"MYTRAVELNOTEBOOK_DB_DBDRIVER" envDefault:"postgres"`
	DBConnectionStr   string `env:"MYTRAVELNOTEBOOK_DB_DBCONNECTIONSTR" envDefault:"postgres://postgres:postgrespw@localhost:55000/postgres?sslmode=disable"`
	DBConnectionDebug bool   `env:"FLAGR_DB_DBCONNECTION_DEBUG" envDefault:"true"`
	// DBConnectionRetryAttempts controls how we are going to retry on db connection when start the flagr server
	DBConnectionRetryAttempts uint          `env:"FLAGR_DB_DBCONNECTION_RETRY_ATTEMPTS" envDefault:"9"`
	DBConnectionRetryDelay    time.Duration `env:"FLAGR_DB_DBCONNECTION_RETRY_DELAY" envDefault:"100ms"`

	// logrus stuff
	// LogrusLevel sets the logrus logging level
	LogrusLevel string `env:"FLAGR_LOGRUS_LEVEL" envDefault:"info"`
	// LogrusFormat sets the logrus logging formatter
	// Possible values: text, json
	LogrusFormat string `env:"FLAGR_LOGRUS_FORMAT" envDefault:"text"`

	// 0auth Authent google
	clientID     string `env:"GOOGLE_CLIENT_ID" envDefault:"foobar"`
	clientSecret string `env:"GOOGLE_CLIENT_SECRET" envDefault:"foobar"`
}{}
