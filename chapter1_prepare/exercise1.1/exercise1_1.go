package main

import "fmt"

// 与えられた整数が2の倍数かどうかを調べるアルゴリズムを示しなさい
func main() {
	var x int
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("x is a multiple of 2")
	} else {
		fmt.Println("x is not a multiple of 2")
	}
}
