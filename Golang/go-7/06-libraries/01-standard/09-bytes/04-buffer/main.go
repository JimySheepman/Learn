package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	Sample("Buffer", Buffer)
	Sample("Write", Write)
	Sample("Read", Read)
	Sample("Reset", Reset)
	Sample("Truncate", Truncate)
	Sample("Next", Next)
	Sample("Grow", Grow)
	Sample("Len", Len)
	Sample("Cap", Cap)
}

func Buffer() {
	var buffer bytes.Buffer

	buffer.Write([]byte("Hello "))
	fmt.Fprintf(&buffer, "Gopher!")
	buffer.WriteTo(os.Stdout)
	fmt.Println()
}

func Write() {
	var buffer bytes.Buffer

	buffer.Write([]byte("Hey"))
	buffer.WriteRune(' ')
	buffer.WriteString("there")
	buffer.WriteByte('!')
	buffer.WriteTo(os.Stdout)
	fmt.Println()
}

func Read() {
	var buffer bytes.Buffer

	buffer.WriteString("Hello there! I am Gopher.")
	fmt.Printf("%q\n", buffer.String())
	n, err := buffer.Read([]byte("Hi"))
	fmt.Println(n, err)
	fmt.Printf("%q\n", buffer.String())

	line, err := buffer.ReadString('!')
	fmt.Printf("%q\n", line)
	fmt.Printf("%q\n", buffer.String())

	r, size, err := buffer.ReadRune()
	fmt.Printf("%q %d\n", r, size)
	fmt.Printf("%q\n", buffer.String())

	err = buffer.UnreadByte()
	fmt.Printf("%q\n", buffer.String())
}

func Reset() {
	var buffer bytes.Buffer

	fmt.Printf("%q - %v\n", buffer.String(), buffer)
	buffer.WriteString("Hi")
	fmt.Printf("%q - %v\n", buffer.String(), buffer)
	buffer.Reset()
	fmt.Printf("%q - %v\n", buffer.String(), buffer)
}

func Truncate() {
	var buffer bytes.Buffer

	buffer.WriteString("Hi there!")
	fmt.Printf("%q\n", buffer)
	buffer.Truncate(5)
	fmt.Printf("%q\n", buffer)

}

func Next() {
	var buffer bytes.Buffer

	buffer.WriteString("Hello there!")
	fmt.Printf("%q\n", buffer)
	fmt.Printf("Next(5) -> %q\n", buffer.Next(5))
	fmt.Printf("Next(5) -> %q\n", buffer.Next(5))
}

func Grow() {
	var buffer bytes.Buffer

	fmt.Printf("Cap(%s) -> %d\n", buffer.String(), buffer.Cap())
	buffer.Grow(100)
	fmt.Printf("Cap(%s) -> %d\n", buffer.String(), buffer.Cap())
}

func Len() {
	var buffer bytes.Buffer

	buffer.WriteString("Hi there!")
	fmt.Printf("Len(%s) -> %d\n", buffer.String(), buffer.Len())
}

func Cap() {
	var buffer bytes.Buffer

	buffer.WriteString("Hi there!")
	fmt.Printf("Cap(%s) -> %d\n", buffer.String(), buffer.Cap())
	buffer.WriteString("What's up?")
	fmt.Printf("Cap(%s) -> %d\n", buffer.String(), buffer.Cap())
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
