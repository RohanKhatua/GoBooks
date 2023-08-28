package helpers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//return the hash of the super secret key
func GetSuperSecret() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Could not load Super Secret")
	}

	hash:= sha256.New()
	hash.Write([]byte(os.Getenv("super_secret")))
	return hex.EncodeToString(hash.Sum(nil))

	// fmt.Println("No errors in loading .env")

	// return os.Getenv("super_secret")
}
