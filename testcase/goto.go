package main

import (
	"fmt"
)

func main_goto() {
	goto first
second:
	fmt.Println("SECOND.")
	goto done
first:
	fmt.Println("FIRST.")
	goto second
done:
	fmt.Println("DONE.")
}
