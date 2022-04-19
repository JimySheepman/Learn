package main

import (
	"fmt"
	"strings"
)

func main() {
	Sample("Copmpare", Copmpare)
	Sample("Contains", Contains)
	Sample("ContainsAny", ContainsAny)
	Sample("ContainsRune", ContainsRune)
	Sample("EqualFold", EqualFold)
	Sample("HasPrefix", HasPrefix)
	Sample("HasSuffix", HasSuffix)
}

func Copmpare() {
	// Compares two string
	// If ordered returns -1
	fmt.Printf("%q - %q => %d\n", "a", "b", strings.Compare("a", "b"))
	// If equals returns 0
	fmt.Printf("%q - %q => %d\n", "a", "a", strings.Compare("a", "a"))
	// If reverse ordered returns 1
	fmt.Printf("%q - %q => %d\n", "b", "a", strings.Compare("b", "a"))
}

func Contains() {
	// Searches substr inside str, if exist case-sensitive, returns true
	str, substr := "This is text", "is"
	result := strings.Contains(str, substr)
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func ContainsAny() {
	// Searches all substr chars, if exist any, returns true
	str, substr := "This is text", "at"
	result := strings.ContainsAny(str, substr)
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func ContainsRune() {
	// Searches rune as string in str, if exist, returns true
	str, substr := "This is text", 97
	result := strings.ContainsRune(str, rune(substr))
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func EqualFold() {
	// Checks equality case insensitive
	str, substr := "Go", "go"
	result := strings.EqualFold(str, substr)
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func HasPrefix() {
	// Checks the str starts with substr case sensitive

	fmt.Printf("%q - %q => %t\n", "Gopher", "Go", strings.HasPrefix("Gopher", "Go"))
	fmt.Printf("%q - %q => %t\n", "Gopher", "er", strings.HasPrefix("Gopher", "er"))
	fmt.Printf("%q - %q => %t\n", "Gopher", "", strings.HasPrefix("Gopher", ""))
}

func HasSuffix() {
	// Checks the str ends with substr case sensitive

	fmt.Printf("%q - %q => %t\n", "Gopher", "Go", strings.HasSuffix("Gopher", "Go"))
	fmt.Printf("%q - %q => %t\n", "Gopher", "er", strings.HasSuffix("Gopher", "er"))
	fmt.Printf("%q - %q => %t\n", "Gopher", "", strings.HasSuffix("Gopher", ""))
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
