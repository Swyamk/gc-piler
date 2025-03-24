package main

import "fmt"

func string_map() {
	replacements := map[string]string{
		"\")":  "\"s)",
		"\";":  "\"s;",
		"\",":  "\"s,",
		"\"}":  "\"s}",
		"\" }": "\"s }",
		"\" )": "\"s )",
		"\":":  "\"s:",
	}
	for k, v := range replacements {
		fmt.Print(k + " " + v + " ")
	}
	fmt.Println()
}
