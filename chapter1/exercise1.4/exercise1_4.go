package main

import "fmt"

// 5個の整数a,b,c,d,eの平均と分散を求めるアルゴリズムを示しなさい
func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	m := a + b + c + d + e
	sq := (a*a + b*b + c*c + d*d + e*e) / 5
	v := sq - m*m

	fmt.Printf("m = %+v\n", m)
	fmt.Printf("v = %+v\n", v)
}
