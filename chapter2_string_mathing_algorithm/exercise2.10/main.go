package main

import "fmt"

// Algorithm2.2を改良して、テキストS中に文字xが何回あらわれたかを求める
// アルゴリズムをつくりなさい
func main() {
	var s, w string
	fmt.Scan(&s, &w)

	var i, j, cnt int
	for i+j < len(s) {
		if w[j] == s[i+j] {
			j++
			if j == len(w) {
				cnt++
				j--
				i++
			}
		} else {
			i++
			j = 0
		}
	}
	fmt.Println(cnt)
}
