package main

import (
	"fmt"
	"strings"
)

func timespace() {
	s := "    hello\t   \n "
	fmt.Println("|" + strings.TrimSpace(s) + "|")
}
