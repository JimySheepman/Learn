package son

import "fmt"

func init() {
	fmt.Println("Son package initialized")
}

type Son struct {
	Name string
}

func (s Son) Data(name string) string {
	s.Name = "Son : " + name
	return s.Name
}
