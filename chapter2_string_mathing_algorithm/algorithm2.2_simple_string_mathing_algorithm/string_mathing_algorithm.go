package main

import "fmt"

// 文字列照合アルゴリズム
func main() {
	var s, w string
	fmt.Scan(&s, &w)

	var i, j int
	for i+j < len(s) {
		fmt.Printf("before i = %+v\n", i)
		fmt.Printf("before j = %+v\n", j)
		if w[j] == s[i+j] {
			fmt.Printf("W %#U starts at byte position %d\n", w[j], i)
			fmt.Printf("S %#U starts at byte position %d\n", s[i+j], i)
			// fmt.Printf("i = %+v\n", i)
			// fmt.Printf("j = %+v\n", j)
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
