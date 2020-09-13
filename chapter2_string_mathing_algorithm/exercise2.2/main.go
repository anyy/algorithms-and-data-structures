package main

import "fmt"

// S = baseball と
// x = p に対するAlgorithm2.1の動作を示しなさい。
func main() {
	s := "baseball"
	x := "p"

	for i, c := range s {
		if string(c) == x {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(len(s))
}
