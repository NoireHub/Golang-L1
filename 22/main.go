package main

import (
	"fmt"
	"math/big"
)

func main() {
	exp:= big.NewInt(21)
	a := big.NewInt(1).Exp(big.NewInt(2),exp,nil)
	b := big.NewInt(1).Exp(big.NewInt(2),exp,nil)
	
	fmt.Println(a.Text(10))
	fmt.Println(a.Text(10))

	c := new(big.Int).Mul(a, b)
	fmt.Println("a * b =", c)

	d := new(big.Int).Div(a, b)
	fmt.Println("a / b =", d)

	e := new(big.Int).Add(a, b)
	fmt.Println("a + b =", e)

	f := new(big.Int).Sub(a, b)
	fmt.Println("a - b =", f)
}
