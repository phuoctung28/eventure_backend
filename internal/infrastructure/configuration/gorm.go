package configuration

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_USERNAME = "postgres"
const DB_PASSWORD = "postgres"
const DB_NAME = "eventure"
const DB_HOST = "localhost"
const DB_PORT = "5432"

var Db *gorm.DB

func ProvideGormDB() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
