package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {

	var dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Beirut", dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
