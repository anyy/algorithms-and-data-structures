package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var i int
	for s[i] == s[len(s)-1-i] && i < len(s)/2 {
		i++
	}
	if i >= len(s)/2 {
		fmt.Printf("%sは回文", s)
	} else {
		fmt.Printf("%sは回文でない", s)
	}
}
