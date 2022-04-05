package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}
