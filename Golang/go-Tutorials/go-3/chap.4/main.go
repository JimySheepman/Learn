package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/fatih/color"
	"github.com/koding/logging"
	// <isim> "github.com/koding/logging" aynı paket ismi ile birden fazla olabilir onu adlandıra biliyoruz.
)

func main() {
	// fmt
	fmt.Println("fmt")
	// math/rand
	fmt.Println("My favorite number is", rand.Intn(10))
	// strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("unit_test", "unit"))
	fmt.Println(strings.HasSuffix("unit_test_rar", "rar"))
	fmt.Println(strings.Index("test", "e"))
	color.Red("Error message")
	logging.Info("Uygulama sonu")
}
