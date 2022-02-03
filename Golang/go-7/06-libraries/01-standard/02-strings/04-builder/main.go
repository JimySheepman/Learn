package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.Write([]byte{',', ' '})
	builder.WriteRune(65)
	builder.WriteByte('!')

	fmt.Printf("Len: %d\n", builder.Len())
	fmt.Printf("Cap: %d\n", builder.Cap())
	fmt.Printf("Value: %q\n", builder.String())

	builder.Reset()

	fmt.Printf("Len: %d\n", builder.Len())
	fmt.Printf("Cap: %d\n", builder.Cap())
	fmt.Printf("Value: %q\n", builder.String())

	builder.WriteString("Hello, Gopher!")
	builder.WriteByte(' ')
	builder.WriteString("How are you?")

	fmt.Printf("Len: %d\n", builder.Len())
	fmt.Printf("Cap: %d\n", builder.Cap())
	fmt.Printf("Value: %q\n", builder.String())

	builder.Grow(100)

	fmt.Printf("Len: %d\n", builder.Len())
	fmt.Printf("Cap: %d\n", builder.Cap())
	fmt.Printf("Value: %q\n", builder.String())
}
