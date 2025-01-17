package utils

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if Config == nil {
		Logger.Fatal("Configuration is not loaded")
	}

	dbUser := Config.GetString(DatabaseUsername)
	dbPassword := Config.GetString(DatabasePassword)
	dbHost := Config.GetString(DatabaseHost)
	dbPort := Config.GetInt(DatabasePort)
	dbName := Config.GetString(DatabaseName)
	dbSslMode := Config.GetString(DatabaseSslMode)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		Logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	DB = database
}
