package main

import "fmt"

// 平面上の三点の座標が与えられたとき、それらの三点を頂点としてもつ
// 三角形の面積を求めるアルゴリズムを示しなさい。
// ただし三点が一直線上にあるとき、面積は0になるものとする
func main() {
	var a, b, c, d, e, f int
	x1 := c - a
	y1 := d - b
	x2 := e - a
	y2 := f - b
	S := (x1*y2 - x2*y1) / 2
	if S < 0 {
		S -= S
	}
	fmt.Printf("S = %+v\n", S)
}
