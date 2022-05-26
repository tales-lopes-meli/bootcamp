package main

import (
	"fmt"
	"time"
)

const (
	MSG_HELLO = "Hello, %s\n"
)

func main() {
	fmt.Println(time.Now().UnixNano())
}
