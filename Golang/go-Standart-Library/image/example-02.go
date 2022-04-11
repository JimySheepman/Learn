package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	_ "image/jpeg"
)

func main() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(Data))
	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)
}
