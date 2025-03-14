package main

import (
    "fmt"
    "testing"
)

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func TestChannels(t *testing.T) {
    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    
    x, y := <-c, <-c
    fmt.Println(x, y, x+y)
}