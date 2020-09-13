package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)

	fmt.Println(createTable(w))
}

// パターン照合テーブル
func createTable(w string) []int {
	t := make([]int, len(w))
	t[0] = -1

	for i := 1; i < len(w); i++ {
		k := i - 1
		for k > 0 && w[i-1] != w[k-1] {
			k = k - 1
		}
		t[i] = k
	}
	return t
}
