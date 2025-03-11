package main

import (
    "fmt"
    "testing"
)

type Rectangle struct {
    width, height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (r *Rectangle) scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

func TestStructMethods(t *testing.T) {
    r := Rectangle{5.0, 3.0}
    fmt.Println(r.area())
    r.scale(2.0)
    fmt.Println(r.area())
}