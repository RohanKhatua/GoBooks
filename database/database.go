package database

import (
	"time"

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

func cleanDatabase (db *gorm.DB) {
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM books")
	db.Exec("DELETE FROM cart_items")
	db.Exec("DELETE FROM purchases")
	db.Exec("DELETE FROM reviews")	
}

func SeedDatabase (db *gorm.DB) {
	// Clean database
	cleanDatabase(db)

	// Create sample users
	users := []models.User{
		{UserName: "user1", Password: "pass1", Role: "user", IsActivated: true},
		{UserName: "user2", Password: "pass2", Role: "user", IsActivated: true},
	}
	for i := range users {
		db.Create(&users[i])
	}

	// Create sample books
	books := []models.Book{
		{Author: "Author 1", Year: 2020, Title: "Book 1", Quantity: 10, Description: "Description 1", Price: 20},
		{Author: "Author 2", Year: 2019, Title: "Book 2", Quantity: 5, Description: "Description 2", Price: 15},
	}
	for i := range books {
		db.Create(&books[i])
	}

	// Create sample cart items
	cartItems := []models.CartItem{
		{UserID: users[0].ID, BookID: books[0].ID},
		{UserID: users[1].ID, BookID: books[1].ID},
	}
	for i := range cartItems {
		db.Create(&cartItems[i])
	}

	// Create sample purchases
	purchases := []models.Purchase{
		{UserID: int(users[0].ID), BookID: books[0].ID, Quantity: 2, PurchaseDate: time.Now()},
		{UserID: int(users[1].ID), BookID: books[1].ID, Quantity: 1, PurchaseDate: time.Now()},
	}
	for i := range purchases {
		db.Create(&purchases[i])
	}

	// Create sample reviews
	reviews := []models.Review{
		{UserID: users[0].ID, BookID: books[0].ID, Rating: 4, Comment: "Good book", ReviewDate: time.Now()},
		{UserID: users[1].ID, BookID: books[1].ID, Rating: 5, Comment: "Excellent read", ReviewDate: time.Now()},
	}
	for i := range reviews {
		db.Create(&reviews[i])
	}

	myLogger := customLogger.NewLogger()

	myLogger.Info("Seeding Complete")
}