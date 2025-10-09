package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	strsA := []string{"hello", "world"}
	strsB := strings.Join(strsA, "=")
	fmt.Println(strsB)

	joinRes := bytes.Join([][]byte{[]byte("hello"), []byte("world")}, []byte("="))
	fmt.Printf("joinRes value: %s\n", joinRes)
	fmt.Printf("joinRes value: %x\n", joinRes)
	fmt.Println(string(joinRes))
}
