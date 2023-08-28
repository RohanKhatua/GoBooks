package database

import (
	"github.com/RohanKhatua/fiber-jwt/customLogger"
	"github.com/RohanKhatua/fiber-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

func GlobalActivationScope(db *gorm.DB) *gorm.DB {
	return db.Where("is_activated = ?", true)
}

var Database DbInstance

func ConnectDb () {
	myLogger := customLogger.NewLogger()
	dsn := "user=myuser password=password dbname=mydb host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Scopes(GlobalActivationScope)

	if err!=nil {
		myLogger.Fatal("Could Not Connect to DB")
		// os.Exit(2)
	}

	myLogger.Info("Connected to DB Successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
	myLogger.Info("Runnning Migrations")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.CartItem{})
	db.AutoMigrate(&models.Purchase{})
	db.AutoMigrate(&models.Review{})

	myLogger.Info("Migrations Complete")

	
	Database = DbInstance{Db: db}

}