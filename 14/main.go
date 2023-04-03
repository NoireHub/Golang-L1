package main

import (
	"fmt"
	"reflect"
)

func typeOf(a interface{}) string {
	aType := reflect.TypeOf(a)

	if reflect.ValueOf(a).Kind() == reflect.Chan {
		return "channel"
	}

	return aType.Name()
}

func main() {
	var a chan bool
	var b int
	var c string
	var d bool

	
	fmt.Println(typeOf(a))
	fmt.Println(typeOf(b))
	fmt.Println(typeOf(c))
	fmt.Println(typeOf(d))
	
}
