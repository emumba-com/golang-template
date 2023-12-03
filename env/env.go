package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type envFile struct {
	BuildEnv   string
	ServerPort string
	DbHost     string
	DbUsername string
	DbPassword string
	DbName     string
	DbPort     string
}

func getEnvValue(key string, defaultValue string) string {
	if envValue := os.Getenv(key); envValue != "" {
		return envValue
	}

	return defaultValue
}

var Env *envFile

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err.Error())
	}

	serverPort := getEnvValue("SERVER_PORT", "8000")

	buildEnv := getEnvValue("BUILD_ENV", "dev")

	dbHost := getEnvValue("DB_HOST", "localhost")
	dbPort := getEnvValue("DB_PORT", "5432")
	dbUsername := getEnvValue("DB_USERNAME", "postgres")
	dbPassword := getEnvValue("DB_PASSWORD", "1234")
	dbName := getEnvValue("DB_NAME", "gin_service")

	Env = &envFile{
		BuildEnv:   buildEnv,
		ServerPort: serverPort,
		DbHost:     dbHost,
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbName:     dbName,
		DbPort:     dbPort,
	}
}
