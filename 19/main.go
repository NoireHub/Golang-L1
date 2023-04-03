package main

import "fmt"

func reverse(str string) string {
	data := []rune(str)

	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }

	return string(data)
}

func main() {
	var data string
	fmt.Print("enter string: ")
	fmt.Scan(&data)
	data = reverse(data)
	fmt.Println(data)
}

// abcdfgry
// 漢字
//	อักษรจีน
// かんじ