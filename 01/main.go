package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	ActionName string
}

func main() {
	action := Action{
		Human{"Valera", 54},
		 "Run",
		}
		
	fmt.Println(action.GetName())
}

func (h *Human) GetName() string {
	return h.Name
}