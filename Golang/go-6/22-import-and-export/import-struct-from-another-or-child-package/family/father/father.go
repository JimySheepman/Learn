package father

import "fmt"

func init() {
	fmt.Println("Father package initialized")
}

type Father struct {
	Name string
}

func (f Father) Data(name string) string {
	f.Name = "Father : " + name
	return f.Name
}
