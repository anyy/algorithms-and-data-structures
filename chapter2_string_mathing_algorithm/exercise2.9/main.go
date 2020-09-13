package main

import "fmt"

// Algorithm2.1を改良して、テキストS中に文字xが何回あらわれたかを求める
// アルゴリズムをつくりなさい
func main() {
	var s, x string
	var cnt int
	fmt.Scan(&s, &x)
	for _, c := range s {
		if string(c) == x {
			cnt++
		}
	}
	fmt.Println(cnt)
}
