package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	ar := constant.MakeFloat64(2.3)
	ai := constant.MakeImag(constant.MakeInt64(5))
	a := constant.BinaryOp(ar, token.ADD, ai)

	b := constant.MakeUint64(11)
	c := constant.BinaryOp(a, token.MUL, b)

	Ar, exact := constant.Float64Val(constant.Real(c))
	if !exact {
		fmt.Printf("Could not represent real part %s exactly as float64\n", constant.Real(c))
	}
	Ai, exact := constant.Float64Val(constant.Imag(c))
	if !exact {
		fmt.Printf("Could not represent imaginary part %s as exactly as float64\n", constant.Imag(c))
	}
	C := complex(Ar, Ai)

	fmt.Println("literal", 25.3+55i)
	fmt.Println("go/constant", c)
	fmt.Println("complex128", C)

}