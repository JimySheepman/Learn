package kata

import "fmt"

func Greet(name string) string {
	if name == "Johnny" {
		name = "my love"
	}
	return fmt.Sprintf("Hello, %v!", name)
}
