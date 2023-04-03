package main

import "fmt"

func main() {
	a:= 12
	b:= 43

	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
