package main

import (
	"fmt"
	"math"
)

// 平面上の二点の座標が与えられたとき、それらの間のユークリッド距離を求める
// アルゴリズムを示しなさい。加減乗除と平方根の各演算は自由に使って良い。
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	x := a - c
	y := b - d
	dist := math.Sqrt(float64(x*x) + float64(y*y))
	fmt.Printf("dist = %+v\n", dist)
}
