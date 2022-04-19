package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://stackoverflow.com/questions/tagged/javascript")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".question-summary .summary").Each(func(i int, s *goquery.Selection) {
		title := s.Find("H3").Text()
		fmt.Println(i, title)
	})
}

type Earthquake struct {
	Timestamp string
	Latitide  string
	Longitude string
	Magnitudo string
	Depth     string
	Territory string
}

var arrayEarthquake []Earthquake

func postScrape() {
	doc, err := goquery.NewDocument("http://www.bmkg.go.id/gempabumi/gempabumi-terkini.bmkg")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody tr").Each(func(_ int, tr *goquery.Selection) {

		e := Earthquake{}

		tr.Find("td").Each(func(ix int, td *goquery.Selection) {
			switch ix {
			case 1:
				e.Timestamp = td.Text()
			case 2:
				e.Latitide = td.Text()
			case 3:
				e.Longitude = td.Text()
			case 4:
				e.Magnitudo = td.Text()
			case 5:
				e.Depth = td.Text()
			case 6:
				e.Territory = td.Text()
			}
		})

		arrayEarthquake = append(arrayEarthquake, e)
	})

	fmt.Println(arrayEarthquake)
}

func main() {
	ExampleScrape()
	postScrape()
}
