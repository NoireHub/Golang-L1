package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	words := strings.Split(str, " ")
	rWords := make([]string, len(words))

	for i, j := len(words)-1, 0; i >= 0; i, j = i-1, j+1 {
		rWords[j] = words[i]
	}

	newS := strings.Join(rWords, " ")
	fmt.Println(newS)
}
