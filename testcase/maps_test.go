package main

import (
    "fmt"
    "testing"
)

func TestMaps(t *testing.T) {
    map1 := make(map[string]int)
    map1["one"] = 1
    map1["two"] = 2
    
    map2 := map[string]string{
        "name": "John",
        "city": "New York",
    }
    
    delete(map1, "one")
    
    if val, exists := map2["name"]; exists {
        fmt.Println("Name exists:", val)
    }
    
    for key, value := range map1 {
        fmt.Println(key, value)
    }
}