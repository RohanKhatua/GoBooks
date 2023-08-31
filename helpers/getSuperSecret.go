package helpers

import (
	"crypto/sha256"
	"encoding/hex"

	"os"


)

//return the hash of the super secret key
func GetSuperSecret() string {
	

	hash:= sha256.New()
	hash.Write([]byte(os.Getenv("super_secret")))
	return hex.EncodeToString(hash.Sum(nil))

	// fmt.Println("No errors in loading .env")

	// return os.Getenv("super_secret")
}
