package config

// Config is the whole configuration of the app
var Config = struct {
	// Host - ARGOSTATUS server host
	Host string `env:"HOST" envDefault:"localhost"`
	// Port - ARGOSTATUS server port
	Port int `env:"PORT" envDefault:"18000"`

	// LogrusLevel sets the logrus logging level
	LogrusLevel string `env:"ARGOSTATUS_LOGRUS_LEVEL" envDefault:"info"`
	// LogrusFormat sets the logrus logging formatter
	// Possible values: text, json
	LogrusFormat string `env:"ARGOSTATUS_LOGRUS_FORMAT" envDefault:"text"`

	// MiddlewareVerboseLoggerEnabled - to enable the negroni-logrus logger for all the endpoints useful for debugging
	MiddlewareVerboseLoggerEnabled bool `env:"ARGOSTATUS_MIDDLEWARE_VERBOSE_LOGGER_ENABLED" envDefault:"true"`
	// MiddlewareVerboseLoggerExcludeURLs - to exclude urls from the verbose logger via comma separated list
	MiddlewareVerboseLoggerExcludeURLs []string `env:"ARGOSTATUS_MIDDLEWARE_VERBOSE_LOGGER_EXCLUDE_URLS" envDefault:"" envSeparator:","`
	// MiddlewareGzipEnabled - to enable gzip middleware
	MiddlewareGzipEnabled bool `env:"ARGOSTATUS_MIDDLEWARE_GZIP_ENABLED" envDefault:"true"`

	// CORSEnabled - enable CORS
	CORSEnabled          bool     `env:"ARGOSTATUS_CORS_ENABLED" envDefault:"true"`
	CORSAllowCredentials bool     `env:"ARGOSTATUS_CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	CORSAllowedHeaders   []string `env:"ARGOSTATUS_CORS_ALLOWED_HEADERS" envDefault:"Origin,Accept,Content-Type,X-Requested-With,Authorization,Time_Zone" envSeparator:","`
	CORSAllowedMethods   []string `env:"ARGOSTATUS_CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,PATCH" envSeparator:","`
	CORSAllowedOrigins   []string `env:"ARGOSTATUS_CORS_ALLOWED_ORIGINS" envDefault:"*" envSeparator:","`
	CORSExposedHeaders   []string `env:"ARGOSTATUS_CORS_EXPOSED_HEADERS" envDefault:"WWW-Authenticate" envSeparator:","`

	// WebPrefix - base path for web and API
	// e.g. ARGOSTATUS_WEB_PREFIX=/foo
	// UI path  => localhost:18000/foo"
	// API path => localhost:18000/foo/api/v1"
	WebPrefix string `env:"ARGOSTATUS_WEB_PREFIX" envDefault:""`
}{}
