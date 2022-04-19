package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	keyFile       = "aes.key"
	encryptedFile = "aes.enc"
)

var IV = []byte("1234567812345678")

func readKey(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, _ := pem.Decode(key)
	return block.Bytes, nil
}

func createKey() []byte {
	genkey := make([]byte, 16)
	_, err := rand.Read(genkey)
	if err != nil {
		log.Fatalf("Failed to read new random key: %s", err)
	}
	return genkey
}

func saveKey(filename string, key []byte) {
	block := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0644)
	if err != nil {
		log.Fatalf("Failed in saving key to %s: %s", filename, err)
	}
}

func aesKey() []byte {
	file := fmt.Sprintf(keyFile)
	key, err := readKey(file)
	if err != nil {
		log.Println("Creating a new AES key")
		key = createKey()
		saveKey(file, key)
	}
	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}

func encryption(plainText string) {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	err := ioutil.WriteFile(fmt.Sprintf(encryptedFile), bytes, 0644)
	if err != nil {
		log.Fatalf("Writing encryption file: %s", err)
	} else {
		fmt.Printf("Message encrypted in file: %s\n\n", encryptedFile)
	}
}

func decryption() []byte {
	bytes, err := ioutil.ReadFile(fmt.Sprintf(encryptedFile))
	if err != nil {
		log.Fatalf("Reading encrypted file: %s", err)
	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}

func main() {

	var plainText = "AES is now being used worldwide for encrypting digital information, including financial, and government data."
	encryption(plainText)

	fmt.Printf("Decrypted Message: %s", decryption())
}
