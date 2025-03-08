package main

import (
    "fmt"
)

func isPalindrome(number int) bool {
    original := number
    reversed := 0

    for number > 0 {
        digit := number % 10
        reversed = reversed*10 + digit
        number /= 10
    }

    return original == reversed
}

func main() {
    number := 121 // You can change this number to test other cases
    if isPalindrome(number) {
        fmt.Printf("%d is a palindrome\n", number)
    } else {
        fmt.Printf("%d is not a palindrome\n", number)
    }
}