package database

import (
	"log"
	"os"

	"github.com/RohanKhatua/fiber-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb () {
	dsn := "user=myuser password=password dbname=mydb host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil {
		log.Fatal("Failed to connect to DB")
		os.Exit(2)
	}

	log.Println("Connected to DB Successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{})

	Database = DbInstance{Db: db}

}