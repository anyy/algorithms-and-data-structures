package main

import "fmt"

// 文字列照合アルゴリズム
func main() {
	var s, w string
	fmt.Scan(&s, &w)

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
