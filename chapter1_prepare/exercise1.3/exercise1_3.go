package main

import "fmt"

// 5個の整数a,b,c,d,eの最大値を求めるアルゴリズムを示しなさい
func main() {
	var a, b, c, d, e, max int
	fmt.Scan(&a, &b, &c, &d, &e)

	max = a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	if e > max {
		max = e
	}

	fmt.Println(max)
}
