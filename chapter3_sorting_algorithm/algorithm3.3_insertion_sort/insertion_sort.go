package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(insertionSort(nums))
}

func insertionSort(n []int) []int {
	for i := 1; i < len(n); i++ {
		j := i
		for j > 0 {
			if n[j-1] > n[j] {
				n[j-1], n[j] = n[j], n[j-1]
			}
			j = j - 1
		}
	}
	return n
}
