package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(gistRequestJson))
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("TOKEN")))

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
}
