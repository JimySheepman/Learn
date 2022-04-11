package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}
