package common

import (
	"fmt"
	"log"
	"os"

	"gostudio_app/customer"
	"gostudio_app/department"
	"gostudio_app/employee"
	"gostudio_app/order"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // Declare a global variable for the database connection

func setupDatabase() *gorm.DB {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.Exec("CREATE SCHEMA IF NOT EXISTS gostudio").Error
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	err = db.Exec("SET search_path TO gostudio").Error
	if err != nil {
		log.Fatalf("failed to set search_path: %v", err)
	}

	err = db.AutoMigrate(
		&customer.Customer{},
		&order.Order{},
		&employee.Employee{},
		&department.Department{},
	)

	if err != nil {
		log.Fatalf("failed to perform database migration: %v", err)
	}

	log.Println("Database migration successful")

	return db

}

func SetupDatabase() *gorm.DB {
	return setupDatabase()
}
