package main

import (
	"fmt"
)

func range_map_value() {
	m := map[string]string{"first": "hi", "second": "you", "third": "there"}
	first := true
	for _, v := range m {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
