package main

import "fmt"

// 鶴亀算をアルゴリズムの形で記述しなさい。
// ただし入力は足の数の合計xと個体の総数yとし、
// 出力はつると亀のそれぞれの個体数とすること。
// また、解が存在しない入力を与えたときには、そのことをきちんと出力させること。
// ちなみに、x=20, y=6の時の具体的な実行の様子は以下の通りである。
// 「6羽すべてが鶴と仮定すると、足の数は6x2=12(ステップ1)」
// しかし実際には足の数は20本だから、全てが鶴と仮定したときとの差は20-12=8本(ステップ2)
// 鶴と亀の足の数の差は1匹・羽当たり4-2=2本であるから(ステップ3)、
// 全体で8本の差が出るということは、亀の数は8/2=4匹、鶴の数は6-4=2羽(ステップ4)」
func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// 鶴の個体数c, 亀の個体数t
	// 仮に全てが鶴とする
	c := y
	// 足の不足分
	d := x - 2*c
	// 必要な亀の数
	t := d / (4 - 2)
	// 鶴の数
	c = y - t
	if c >= 0 && t >= 0 {
		fmt.Printf("鶴 = %+v匹\n", c)
		fmt.Printf("亀 = %+v匹\n", t)
	}
}
