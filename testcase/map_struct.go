package main

import (
	"fmt"
)

type _Creature struct {
	X    float64
	Y    float64
	name string
}

func map_struct() {
	o := map[int]*Creature{
		0: &Creature{1.2, 3.4, "Bob"},
		1: &Creature{4.5, 6.7, "Alice"},
	}
	fmt.Println(o[1])
}
