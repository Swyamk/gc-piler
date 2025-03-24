package main

import (
	"fmt"
)

func for_range_single() {
	l := []string{"a", "b", "c"}
	for i := range l {
		fmt.Println(i)
	}
}
