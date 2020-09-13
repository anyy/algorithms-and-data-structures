package main

import "fmt"

// ある整数の2乗であらわされる整数を平方数と呼ぶ。
// 例えば4, 49, 169などはいずれも平方数である。
// 与えられた自然数xが平方数かどうかを調べるアルゴリズムを示しなさい
func main() {
	var x int
	fmt.Scan(&x)

	var i int
	for i <= x && x != i*i {
		i = i + 1
	}

	if i <= x {
		fmt.Printf("i = %+v\n", i)
	}
}
