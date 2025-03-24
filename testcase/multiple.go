package main

import (
	"fmt"
)

func addsub(x int) (a, b int) {
	return x + 2, x - 2
}

func multiple() {
	y, z := addsub(4)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
}
