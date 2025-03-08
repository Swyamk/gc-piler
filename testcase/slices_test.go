package main

import (
    "fmt"
    "testing"
)

func TestSlices(t *testing.T) {
    slice1 := make([]int, 3, 5)
    slice2 := []int{1, 2, 3, 4, 5}

    slice2 = append(slice2, 6)
    slice3 := slice2[1:4]

    fmt.Println(slice1, len(slice1), cap(slice1))
    fmt.Println(slice2, len(slice2), cap(slice2))
    fmt.Println(slice3, len(slice3), cap(slice3))
}