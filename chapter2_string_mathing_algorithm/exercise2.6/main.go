package main

import "fmt"

// S = baseball と
// x = bat に対するAlgorithm2.3の動作を示しなさい。
func main() {
	s := "baseball"
	w := "bat"

	t := createTable(w)
	var i, j int
	for i+j < len(s) {
		if w[j] == s[i+j] {
			j++
			if j == len(w) {
				fmt.Println(i)
				return
			}
		} else {
			i = i + j - t[j]
			if j > 0 {
				j = t[j]
			}
		}
	}
	fmt.Println(len(s))
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
