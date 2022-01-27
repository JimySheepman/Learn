package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

var secretKey = "4234kxzjcjj3nxnxbcvsjfj"

// Generate a salt string with 16 bytes of crypto/rand data.
func generateSalt() string {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func main() {
	message := "Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model."
	salt := generateSalt()
	fmt.Println("Message: " + message)
	fmt.Println("\nSalt: " + salt)

	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\nHMAC-Sha256: %x", hash.Sum(nil))

	hash = hmac.New(sha512.New, []byte(secretKey))
	io.WriteString(hash, message+salt)
	fmt.Printf("\n\nHMAC-sha512: %x", hash.Sum(nil))
}
