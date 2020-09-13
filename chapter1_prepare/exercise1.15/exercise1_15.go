package main

import "fmt"

// 我々が使っているコンピュータでは、整数としてあらわされる数には上限があり、
// その上限を超えると桁あふれ(overflow)が発生し正しい計算が行えない。
// 桁あふれが起こらないように配慮しつつ、n番目のカタラン数を計算するアルゴリズムを作りなさい。
// n番目のカタラン数は、次のように定義される。
func main() {
	var n int
	fmt.Scan(&n)

	c := 1
	for i := 1; i <= n; i++ {
		c = c * (n + i) / i
	}
	fmt.Printf("Catalan number = %+v\n", c/(n+1))
}
