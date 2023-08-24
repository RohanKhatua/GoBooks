package helpers

import (
	"github.com/joho/godotenv"
	"os"
	"log"
)

func GetSuperSecret() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load Super Secret")
	}

	// fmt.Println("No errors in loading .env")

	return os.Getenv("super_secret")
}
