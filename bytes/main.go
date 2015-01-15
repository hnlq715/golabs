package main

import (
	"bytes"
	"fmt"
)

// bytes.Buffer's init cap is 64
func main() {
	var buf bytes.Buffer
	buf.WriteString("Hello world")
	fmt.Println(buf.Len(), cap(buf.Bytes()))
}
