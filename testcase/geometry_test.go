package main

import (
    "fmt"
    "testing"
)

type Geometry interface {
    area() float64
    perim() float64
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return 3.14159 * c.radius * c.radius
}

func (c Circle) perim() float64 {
    return 2 * 3.14159 * c.radius
}

func measure(g Geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func TestGeometry(t *testing.T) {
    c := Circle{5}
    measure(c)
}