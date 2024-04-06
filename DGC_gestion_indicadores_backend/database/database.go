package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	var err error
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, username, password, dbname, port)
	log.Println("dsn : ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
		return nil
	}
	log.Println("`Successfully connected to the database")

	return db
}
