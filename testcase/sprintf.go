package main

import (
	"fmt"
)

var globalMessage string

func sprintf() {
	globalMessage = fmt.Sprintf("hello: %v != %v", 123, 256)
	fmt.Println(globalMessage)
}
