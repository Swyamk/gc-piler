package main

import (
	"fmt"
)

func for_regular() {
	for i := 0; i < 10; i++ {
		fmt.Println("UP   :", i)
	}
	x := 42
	for ; x > 0; x -= 10 {
		fmt.Println("DOWN :", x)
	}
}
