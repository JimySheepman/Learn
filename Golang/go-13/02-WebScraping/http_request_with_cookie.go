package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	myCookie := &http.Cookie{
		Name:  "cookieKey1",
		Value: "value1",
	}

	request.AddCookie(myCookie)

	fmt.Println(request.Cookies())
	fmt.Println(request.Header)
}
