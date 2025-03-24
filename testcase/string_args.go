package main

import "fmt"

func threecat(a, b, c string) string {
	return a + b + c
}

func string_args() {
	threecat("in", "cred", "ible")
	fmt.Println("hi")
}
