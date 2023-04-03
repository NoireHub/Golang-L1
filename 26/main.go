package main

import (
	"fmt"
	"unicode"
)

func isUnique(s string) bool {
	uniqueMap := make(map[rune]bool)

	for _, val := range s {
		if !uniqueMap[unicode.ToLower(val)]{
			uniqueMap[unicode.ToLower(val)] = true		
		}else{
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isUnique("漢字漢字"))
	fmt.Println(isUnique("abcdfgry"))

}

// abcdfgry
// 漢字
//	อักษรจีน
// かんじ