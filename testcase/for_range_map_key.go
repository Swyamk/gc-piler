package main 

import(
	"fmt"
)

func for_range_map_key() {
	m := map[string]string{"first": "hi", "second": "you", "third": "there"}
	first := true
	for k := range m {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(k)
	}
	fmt.Println()
}