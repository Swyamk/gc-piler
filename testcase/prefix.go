package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("asdfqwerty", "asdf"))
	fmt.Println(strings.HasPrefix("asdfqwerty", "qwerty"))
}
