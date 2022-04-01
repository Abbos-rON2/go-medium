package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App           string
	Environment   string // development, staging, production
	LogLevel      string // debug, info, warn, error, dpanic, panic, fatal
	ServiceScheme string
	ServiceHost   string
	ServicePort   string

	HTTPHost string
	HTTPPort string

	BasePath string

	DBHost    string
	DBPort    string
	DBName    string
	DBUser    string
	DBPass    string
	JwtSecret string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "anor_admin_api_gateway"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))

	config.ServiceScheme = cast.ToString(getOrReturnDefaultValue("SERVICE_SCHEME", "http"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":8001"))

	config.HTTPHost = cast.ToString(getOrReturnDefaultValue("HTTP_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8001"))
	config.BasePath = cast.ToString(getOrReturnDefaultValue("BASE_PATH", "/v1"))
	config.DBHost = cast.ToString(getOrReturnDefaultValue("DB_HOST", "localhost"))
	config.DBPort = cast.ToString(getOrReturnDefaultValue("DB_PORT", "5432"))
	config.DBName = cast.ToString(getOrReturnDefaultValue("DB_NAME", "postgres"))
	config.DBUser = cast.ToString(getOrReturnDefaultValue("DB_USER", "postgres"))
	config.DBPass = cast.ToString(getOrReturnDefaultValue("DB_PASS", "postgres"))
	config.JwtSecret = cast.ToString(getOrReturnDefaultValue("JWT_SECRET", "secret"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
