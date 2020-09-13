package main

import "fmt"

// 入力としてテキストSと文字xが与えられたとき、
// S中でxが最初にあらわれる位置の添字を求めるアルゴリズムを示せ。
// ただしxがSにあらわれないときは、質問に対する答えとして、
// 添字|S|を返すものとする。
// Sの添字は0以上|S|-1以下であるから、|S|が返されることで、
// Sの中にxがあらわれないことが正しく判断できることに注意しよう。
func main() {
	var s, x string
	fmt.Scan(&s, &x)
	for i, c := range s {
		if string(c) == x {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(len(s))
}
