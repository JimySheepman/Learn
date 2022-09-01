package main

import "fmt"

type testScore struct {
	studentName string
	score       uint8
}

func main() {
	results := []testScore{
		{"John Doe", 20},
		{"Patrick Indexing", 15},
		{"Bob Ferring", 7},
		{"Claire Novalingua", 8},
	}
	fmt.Println(results)

	for _, result := range results {
		if result.studentName == "Claire Novalingua" {
			fmt.Println("Score Found:", result.score)
		}
	}

	m := make(map[string]int)
	_ = m

	worldCupWinners := map[int]string{
		1930: "Uruguay",
		1934: "Italy",
		1938: "Italy",
		1950: "Uruguay"}
	fmt.Println(worldCupWinners)

}
