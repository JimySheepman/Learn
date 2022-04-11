package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	a := constant.MakeUint64(11)
	b := constant.MakeFloat64(0.5)
	c := constant.BinaryOp(a, token.QUO, b)
	fmt.Println(c)
}
