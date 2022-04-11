package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	const input = `
package temperature
import "fmt"
type Celsius float64
func (c Celsius) String() string  { return fmt.Sprintf("%g°C", c) }
func (c *Celsius) SetF(f float64) { *c = Celsius(f - 32 / 9 * 5) }

type S struct { I; m int }
type I interface { m() byte }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "celsius.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	celsius := pkg.Scope().Lookup("Celsius").Type()
	for _, t := range []types.Type{celsius, types.NewPointer(celsius)} {
		fmt.Printf("Method set of %s:\n", t)
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			fmt.Println(mset.At(i))
		}
		fmt.Println()
	}

	styp := pkg.Scope().Lookup("S").Type()
	fmt.Printf("Method set of %s:\n", styp)
	fmt.Println(types.NewMethodSet(styp))
}
