package main

import "fmt"

// 二つの整数a,bの最小公倍数を求めるアルゴリズムを示しなさい。
// 小学生のときに教わった方法でかまわないが、図的にではなく、プログラム的に記述しなさい
func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := gcd(a, b)
	m := a * b / c
	fmt.Printf("m = %+v\n", m)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
