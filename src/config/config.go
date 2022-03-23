package config

import (
	"fmt"
	"os"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/router"
	"github.com/RaphaSalomao/gin-budget-control/security"
	"github.com/joho/godotenv"
)

func LoadEnvironment(path ...string) {
	godotenv.Load(path...)

	database.DbHost = os.Getenv("DB_HOST")
	database.DbUser = os.Getenv("DB_USER")
	database.DbName = os.Getenv("DB_NAME")
	database.DbPort = os.Getenv("DB_PORT")
	database.DbSslMode = os.Getenv("DB_SSLMODE")
	database.DbPassword = os.Getenv("DB_PASSWORD")
	database.MigrationSourcePath = os.Getenv("M_PATH")

	router.SrvPort = fmt.Sprintf(":%s", os.Getenv("SRV_PORT"))
	
	security.SecretKey = []byte(os.Getenv("JWT_SECRET"))
}
