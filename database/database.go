package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/delta/DAuth-backend-v2/config"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(config config.Config) *gorm.DB {
	username := config.Get("DB_USERNAME")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	dbName := config.Get("DB_NAME")

	logFile, err := os.OpenFile("./logs/database.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(color.RedString("Error opening logs: %s", err))
	}

	loggerDb := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: loggerDb,
	})

	if err != nil {
		panic(color.RedString("Error: %s", err))
	} else {
		fmt.Printf(color.GreenString("Database connected!"))
	}

	return db
}
