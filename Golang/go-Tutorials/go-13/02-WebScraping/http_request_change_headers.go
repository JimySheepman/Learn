package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
