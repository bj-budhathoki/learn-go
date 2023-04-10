package main

import (
	"bytes"
	"fmt"
)

// belong to bytes package of golang
// variable-sized buffer of bytes with read and write methods
// use to manipulate string
func main() {
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Hello")
	strBuffer.WriteString("World")
	fmt.Println("The string buffer output is", strBuffer.String())

}
