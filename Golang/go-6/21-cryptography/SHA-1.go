package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println("\n----------------Small Message----------------")
	message := []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n\n----------------Large Message----------------")
	message = []byte("Today web engineering has modern apps adhere to what is known as a single-page app (SPA) model. This model gives you an experience in which you never navigate to particular pages or even reload a page.  It loads and unloads the various views of our app into the same page itself. If you've ever run popular web apps like Gmail, Facebook, Instagram, or Twitter, you've used a single-page app. In all those apps, the content gets dynamically displayed without requiring you to refresh or navigate to a different page. React gives you a powerful subjective model to work with and supports you to build user interfaces in a declarative and component-driven way.")

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}
