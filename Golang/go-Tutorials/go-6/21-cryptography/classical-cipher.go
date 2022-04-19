package main

import (
	"fmt"
	"unicode"
)

// Cipher encrypts and decrypts a string.
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

// Cipher holds the key used to encrypts and decrypts messages.
type cipher []int

// cipherAlgorithm encodes a letter based on some function.
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

// Encryption encrypts a message.
func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// Decryption decrypts a message.
func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

// NewCaesar creates a new Caesar shift cipher.
func NewCaesar(key int) Cipher {
	return NewShift(key)
}

// NewShift creates a new Shift cipher.
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipher([]int{shift})
	return &c
}

func main() {
	c := NewCaesar(1)
	fmt.Println("Encrypt Key(01) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(01) bcde =>", c.Decryption("bcde"))
	fmt.Println()

	c = NewCaesar(10)
	fmt.Println("Encrypt Key(10) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(10) klmn =>", c.Decryption("klmn"))
	fmt.Println()

	c = NewCaesar(15)
	fmt.Println("Encrypt Key(15) abcd =>", c.Encryption("abcd"))
	fmt.Println("Decrypt Key(15) pqrs =>", c.Decryption("pqrs"))
}
