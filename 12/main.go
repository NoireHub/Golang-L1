package main

import "fmt"

type void struct{}

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]void)

	for _, value := range data {
		set[value] = void{}
	}

	for key := range set {
		fmt.Printf("%s ", key)
	}
}
