package main

import "fmt"

// Algorithm2.4を使ってW = aabaaba
// に対するパターン照合テーブルをつくりなさい。
func main() {
	fmt.Println(createTable("aabaaba"))
}

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
