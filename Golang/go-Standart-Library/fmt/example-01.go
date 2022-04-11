package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	integer := 23
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	fmt.Printf("%T %T\n", integer, &integer)

	truth := true
	fmt.Printf("%v %t\n", truth, truth)

	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)

	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)

	point := 110.7 + 25.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)

	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)

	placeholder := `foo bar`
	fmt.Printf("%v %s %q %#q\n", placeholder, placeholder, placeholder, placeholder)

	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)

	isLegume := map[string]bool{
		"peanut ":   true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)

	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	fmt.Printf("%v %p\n", pointer, pointer)

	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)

	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)

	now := time.Unix(123456789, 0).UTC()
	fmt.Printf("%v %q\n", now, now)
}
