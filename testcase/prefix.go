package main

import (
	"fmt"
	"strings"
)

func prefix() {
	fmt.Println(strings.HasPrefix("asdfqwerty", "asdf"))
	fmt.Println(strings.HasPrefix("asdfqwerty", "qwerty"))
}
