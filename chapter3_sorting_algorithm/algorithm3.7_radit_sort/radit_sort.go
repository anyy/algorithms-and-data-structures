package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	raditSort(nums, len(nums), 100)
	fmt.Println(nums)
}

func raditSort(x []int, n, r int) { // n:データ数, r:基数を取り出す最大値
	m := 1                //基数を取り出す桁
	rad := make([]int, n) //基数をしまう配列
	y := make([]int, n)   //作業用配列
	for m <= r {
		for i := 0; i < n; i++ {
			rad[i] = (x[i] / m) % 10 //基数を取り出しrad[i]に保存
		}
		k := 0
		for i := 0; i <= 9; i++ { //基数は 0 から 9 まで
			for j := 0; j < n; j++ {
				if rad[j] == i { //基数の小さいものからy[ ]にコピー
					y[k] = x[j]
					k++
				}
			}
		}
		for i := 0; i < n; i++ {
			x[i] = y[i]
		}
		m *= 10
		fmt.Println(x)
	}
}
