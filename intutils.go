package main

import (
	"strconv"
)

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	isFloat := (err == nil)
	_, err = strconv.ParseInt(s, 0, 64)
	isInt := (err == nil)
	return isFloat || isInt
}

func trailingNumber(s string) int {
	runes := []rune(s)
	lastIndex := len(runes) - 1
	var num string
TRAILINGNUMBEROUT:
	for i := lastIndex; i >= 0; i-- {
		switch runes[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num += string(runes[i]) + num
		default:
			break TRAILINGNUMBEROUT
		}
	}

	n, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return n
}
