package currency

import "fmt"

var f = func() string {
	fmt.Println("variable f initialized")
	return "test"
}()

func init() {
	fmt.Println("currency init")
}

func EuroSymbol() string {
	return "EUR"
}
