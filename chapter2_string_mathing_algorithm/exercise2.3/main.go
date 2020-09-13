package main

import "fmt"

// S = baseball と
// x = ball に対するAlgorithm2.2の動作を示しなさい。
func main() {
	s := "baseball"
	w := "ball"

	var i, j int
	for i+j < len(s) {
		if w[j] == s[i+j] {
			j++
			if j == len(w) {
				fmt.Println(i)
				return
			}
		} else {
			i++
			j = 0
		}
	}
	fmt.Println(len(s))
}
