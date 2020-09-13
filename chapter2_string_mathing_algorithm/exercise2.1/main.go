package main

import "fmt"

// S = baseball と
// x = e に対するAlgorithm2.1の動作を示しなさい。
func main() {
	s := "baseball"
	x := "e"

	for i, c := range s {
		if string(c) == x {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(len(s))
}
