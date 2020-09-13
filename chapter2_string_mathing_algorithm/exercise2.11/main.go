package main

import "fmt"

// 文字列 gateman を反転させると、
// nametagという文字列になる。
// 入力された文字列からその反転文字列を生成するアルゴリズムをつくりなさい。
func main() {
	s := "gateman"
	var r string
	for i := 0; i < len(s); i++ {
		r += string(s[len(s)-1-i])
	}
	fmt.Printf("r = %+v\n", r)
}
