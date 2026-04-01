package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	AppPort string

	DBDriver string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string

	PasetoKey string

	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 7 * 24 * time.Hour
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file")
	}

	AppPort = os.Getenv("APP_PORT")

	DBDriver = os.Getenv("DB_DRIVER")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")

	PasetoKey = os.Getenv("PASETO_KEY")
}