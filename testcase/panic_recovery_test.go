package main

import (
    "fmt"
    "testing"
)

func recoverExample() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from:", r)
    }
}

func mayPanic() {
    defer recoverExample()
    panic("something went wrong")
}

func TestPanicRecover(t *testing.T) {
    mayPanic()
    fmt.Println("After panic")
}