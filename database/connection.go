package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ario-team/glassnode-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Viper.GetString("DB_HOST"), config.Viper.GetString("DB_USER"), config.Viper.GetString("DB_PASS"), config.Viper.GetString("DB_NAME"), config.Viper.GetString("DB_PORT"))
	log := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		panic(err)
	}
	Connection = db
}
