package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(selectionSort(nums))
}

func selectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		idx, _ := min(n[i:len(n)])
		n[i], n[i+idx] = n[i+idx], n[i]
	}
	return n
}

func min(a []int) (idx, n int) {
	n = a[0]
	idx = 0
	for i, v := range a {
		if n > v {
			n = v
			idx = i
		}
	}
	return
}
