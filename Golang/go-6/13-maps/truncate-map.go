package main

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	// Method - I
	for k := range employee {
		delete(employee, k)
	}

	// Method - II
	employee = make(map[string]int)
}
