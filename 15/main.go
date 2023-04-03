package main


import (
	"bytes"
	"fmt"
)

func someFunc() string {
	var buffer bytes.Buffer
	defer buffer.Reset()
	
	v := createHugeString(1 << 10)
	
	buffer.WriteString(v)
	
	return buffer.String()
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}

func createHugeString(size uint64) string {
	return string(make([]rune, size))
}